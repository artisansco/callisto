import adapter from "@sveltejs/adapter-node";

const config = {
  kit: {
    adapter: adapter(),
    experimental: {
      remoteFunctions: true
    }
  },
  compilerOptions: {
    experimental: {
      async: true
    }
  }
};

export default config;
