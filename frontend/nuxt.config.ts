// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  router: {
    options: {
      hashMode: true,
    },
  },
  hooks: {
    "prerender:routes"({ routes }) {
      routes.clear();
    },
  },
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      versionBuild: "",
    },
  },
  modules: ["@nuxt/ui", "@pinia/nuxt", "@vueuse/nuxt"],
  icon: {
    serverBundle: false,
    clientBundle: {
      scan: true,
    },
  },
  colorMode: {
    preference: "dark",
  },
  css: ["~/assets/css/main.css"],
  nitro: {
    devProxy: {
      "/api": {
        target: "http://localhost:41991/api",
        changeOrigin: true,
      },
    },
  },
});
