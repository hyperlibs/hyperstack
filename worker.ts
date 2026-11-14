export interface Env {
  ASSETS: {
    fetch: (request: Request) => Promise<Response>;
  };
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);

    // Normalize paths without trailing slash for directory routes so index.html resolves correctly
    if (!url.pathname.endsWith("/") && !url.pathname.includes(".")) {
      const slashUrl = new URL(request.url);
      slashUrl.pathname += "/";
      return Response.redirect(slashUrl.toString(), 301);
    }

    return env.ASSETS.fetch(request);
  }
};
