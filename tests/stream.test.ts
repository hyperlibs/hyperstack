import { describe, it, expect, beforeEach } from 'bun:test';
import { ensureObjID, getOrCreateStream, applyStreamPatch, applyLayerMatrix, HyStream } from '../src/htmx-bolt';
import type { StreamPatch } from '../src/types';

class MockElement {
  tagName: string;
  id = '';
  attributes: Record<string, string> = {};
  style: Record<string, string> = {};
  parentElement: MockElement | null = null;
  children: MockElement[] = [];
  textContent = '';
  listeners: Record<string, Function[]> = {};

  constructor(tagName = 'div') {
    this.tagName = tagName.toUpperCase();
  }

  getAttribute(name: string) { return this.attributes[name] !== undefined ? this.attributes[name] : null; }
  setAttribute(name: string, val: string) { this.attributes[name] = String(val); }
  removeAttribute(name: string) { delete this.attributes[name]; }
  hasAttribute(name: string) { return name in this.attributes; }

  appendChild(child: MockElement) {
    child.parentElement = this;
    this.children.push(child);
  }

  closest(selector: string): MockElement | null {
    let curr: MockElement | null = this;
    while (curr) {
      if (selector === '[objid]' && curr.hasAttribute('objid')) return curr;
      curr = curr.parentElement;
    }
    return null;
  }

  querySelectorAll(selector: string): MockElement[] {
    const matched: MockElement[] = [];
    const check = (node: MockElement) => {
      if (selector.includes('.handle') && (node.attributes.class || '').includes('handle')) {
        matched.push(node);
      }
      node.children.forEach(check);
    };
    this.children.forEach(check);
    return matched;
  }

  addEventListener(evt: string, fn: Function) {
    if (!this.listeners[evt]) this.listeners[evt] = [];
    this.listeners[evt].push(fn);
  }

  dispatchEvent(event: any) {
    const list = this.listeners[event.type] || [];
    list.forEach(fn => fn(event));
    if (event.bubbles && this.parentElement) {
      this.parentElement.dispatchEvent(event);
    }
    return true;
  }
}

describe('Hyperstack 1.0 — Universal Addressing (`objid`), `hx-stream` & Layer Matrix', () => {
  let docElements: Map<string, MockElement>;
  let mockDoc: any;

  beforeEach(() => {
    docElements = new Map<string, MockElement>();
    const mockBody = new MockElement('body');
    mockBody.setAttribute('docid', 'doc_main');

    mockDoc = {
      body: mockBody,
      querySelector: (sel: string) => {
        const m = sel.match(/\[objid="([^"]+)"\]/);
        if (m) return docElements.get(m[1]) || null;
        return null;
      },
      querySelectorAll: (sel: string) => {
        return Array.from(docElements.values());
      },
      addEventListener: (evt: string, fn: Function) => mockBody.addEventListener(evt, fn),
      dispatchEvent: (evt: any) => mockBody.dispatchEvent(evt)
    };

    (global as any).document = mockDoc;
    (global as any).window = { CustomEvent: class { type: string; detail: any; bubbles: boolean; constructor(t: string, o: any) { this.type = t; this.detail = o?.detail; this.bubbles = !!o?.bubbles; } } };
    (global as any).CustomEvent = (global as any).window.CustomEvent;
  });

  it('1. Universal Addressing: ensures every element receives a deterministic objid', () => {
    const el = new MockElement('div') as any;
    el.id = 'hero-section';
    expect(el.getAttribute('objid')).toBeNull();

    const id = ensureObjID(el);
    expect(id).toBe('hero-section');
    expect(el.getAttribute('objid')).toBe('hero-section');

    const unnamed = new MockElement('button') as any;
    const autoId = ensureObjID(unnamed);
    expect(autoId).toMatch(/^obj_\d+$/);
    expect(unnamed.getAttribute('objid')).toBe(autoId);
  });

  it('2. Declarative hx-stream: captures invertible deltas and mutates DOM directly', () => {
    const btn = new MockElement('button') as any;
    btn.setAttribute('objid', 'btn_submit');
    btn.textContent = 'Submit';
    btn.style.color = 'blue';
    docElements.set('btn_submit', btn);

    const stream = getOrCreateStream('btn_submit');
    expect(stream.past.length).toBe(0);

    // Apply patch: change text
    const patch: StreamPatch = {
      sub: 'agent_writer',
      dom: 'doc_main',
      target: 'btn_submit',
      op: 'set_text',
      val: 'Confirm Order'
    };

    applyStreamPatch(patch);
    expect(btn.textContent).toBe('Confirm Order');
    expect(stream.past.length).toBe(1);
    expect(stream.past[0].inverse).toBe('Submit');

    // Undo reverts text to original
    stream.undo();
    expect(btn.textContent).toBe('Submit');
    expect(stream.future.length).toBe(1);

    // Redo reapplies forward value
    stream.redo();
    expect(btn.textContent).toBe('Confirm Order');
  });

  it('3. Scoped Undo vs. Global Undo: selected object undos locally without breaking siblings', () => {
    const cardA = new MockElement('div') as any;
    cardA.setAttribute('objid', 'card_a');
    cardA.style.width = '100px';

    const cardB = new MockElement('div') as any;
    cardB.setAttribute('objid', 'card_b');
    cardB.style.color = 'black';

    docElements.set('card_a', cardA);
    docElements.set('card_b', cardB);

    const streamA = getOrCreateStream('card_a');
    const streamB = getOrCreateStream('card_b');

    // Mutate A, then B, then A
    applyStreamPatch({ target: 'card_a', op: 'set_style', prop: 'width', val: '200px' });
    applyStreamPatch({ target: 'card_b', op: 'set_style', prop: 'color', val: 'red' });
    applyStreamPatch({ target: 'card_a', op: 'set_style', prop: 'width', val: '300px' });

    expect(cardA.style.width).toBe('300px');
    expect(cardB.style.color).toBe('red');

    // Scoped Undo on card_b ONLY
    streamB.undo();
    expect(cardB.style.color).toBe('black');
    // Card A is 100% untouched!
    expect(cardA.style.width).toBe('300px');

    // Scoped Undo on card_a
    streamA.undo();
    expect(cardA.style.width).toBe('200px');
    expect(cardB.style.color).toBe('black');
  });

  it('4. Author-Aware Collaborative Undo: undos only requesting author edits', () => {
    const box = new MockElement('div') as any;
    box.setAttribute('objid', 'box_1');
    box.style.opacity = '1';
    docElements.set('box_1', box);

    const stream = getOrCreateStream('box_1');

    // Alice changes opacity, Bob changes opacity
    applyStreamPatch({ sub: 'user_alice', target: 'box_1', op: 'set_style', prop: 'opacity', val: '0.8' });
    applyStreamPatch({ sub: 'user_bob', target: 'box_1', op: 'set_style', prop: 'opacity', val: '0.5' });

    // Alice hits undo on her own changes
    const undone = stream.undo('user_alice');
    expect(undone?.sub).toBe('user_alice');
    expect(undone?.val).toBe('0.8');
  });

  it('5. Two-Phase Lease Model: FocusMode vs EditMode with parent lock bubbling', () => {
    const container = new MockElement('div') as any;
    container.setAttribute('objid', 'section_pricing');

    const button = new MockElement('button') as any;
    button.setAttribute('objid', 'btn_tier');
    container.appendChild(button);

    docElements.set('section_pricing', container);
    docElements.set('btn_tier', button);

    const stream = getOrCreateStream('btn_tier');

    // 1. FocusMode allows sharing
    expect(stream.acquireLease('user_alice', 'focus')).toBe(true);
    expect(button.getAttribute('data-lease')).toBe('focus');

    // 2. EditMode locks the node and bubbles lock to parent container
    expect(stream.acquireLease('user_alice', 'edit')).toBe(true);
    expect(button.getAttribute('data-lease')).toBe('edit');
    expect(container.getAttribute('data-parent-locked')).toBe('true');

    // 3. Another user cannot acquire lease while in EditMode
    expect(stream.acquireLease('user_bob', 'edit')).toBe(false);

    // 4. Release lease unblocks node and parent
    stream.releaseLease('user_alice');
    expect(button.getAttribute('data-lease')).toBeNull();
    expect(container.getAttribute('data-parent-locked')).toBeNull();
  });

  it('6. Casbin & Authorization Context: stream emits canonical tuple (sub, dom, obj, act)', () => {
    const header = new MockElement('h1') as any;
    header.setAttribute('objid', 'main_heading');
    mockDoc.body.appendChild(header);
    docElements.set('main_heading', header);

    let emittedContext: any = null;
    mockDoc.body.addEventListener('hx-stream:patch', (e: any) => {
      emittedContext = e.detail.context;
    });

    applyStreamPatch({
      txId: 'tx_99',
      sub: 'agent_seo',
      dom: 'site_production',
      target: 'main_heading',
      op: 'set_text',
      val: 'AI-Native Web Systems'
    });

    expect(emittedContext).not.toBeNull();
    expect(emittedContext.sub).toBe('agent_seo');
    expect(emittedContext.dom).toBe('site_production');
    expect(emittedContext.obj).toBe('main_heading');
    expect(emittedContext.act).toBe('set_text');
    expect(emittedContext.txId).toBe('tx_99');
  });

  it('7. 2D Layer Matrix: elevation [0..9] and semantic plane positioning', () => {
    const surfaceLayer = new MockElement('div') as any;
    surfaceLayer.setAttribute('level', '2');
    surfaceLayer.setAttribute('role', 'surface');

    const contentLayer = new MockElement('div') as any;
    contentLayer.setAttribute('level', '5');
    contentLayer.setAttribute('role', 'content');

    const manipulatorLayer = new MockElement('div') as any;
    manipulatorLayer.setAttribute('level', '8');
    manipulatorLayer.setAttribute('role', 'manipulator');
    manipulatorLayer.setAttribute('sub', 'handles');

    const handle = new MockElement('span') as any;
    handle.setAttribute('class', 'handle');
    manipulatorLayer.appendChild(handle);

    const hudLayer = new MockElement('div') as any;
    hudLayer.setAttribute('level', '9');
    hudLayer.setAttribute('role', 'hud');

    applyLayerMatrix(surfaceLayer);
    applyLayerMatrix(contentLayer);
    applyLayerMatrix(manipulatorLayer);
    applyLayerMatrix(hudLayer);

    // Surface: level 2 -> z-index 20
    expect(surfaceLayer.style.zIndex).toBe('20');
    // Content: level 5 -> z-index 150
    expect(contentLayer.style.zIndex).toBe('150');
    // Manipulator: level 8 -> z-index 280, pointer-events: none on layer, auto on handle
    expect(manipulatorLayer.style.zIndex).toBe('280');
    expect(manipulatorLayer.style.pointerEvents).toBe('none');
    expect(handle.style.pointerEvents).toBe('auto');
    // HUD: level 9 -> z-index 390, pointer-events: auto
    expect(hudLayer.style.zIndex).toBe('390');
    expect(hudLayer.style.pointerEvents).toBe('auto');
  });
});
