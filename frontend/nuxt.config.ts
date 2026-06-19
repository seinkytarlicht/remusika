// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      versionBuild: "",
    },
  },
  hooks: {
    "prerender:routes"({ routes }) {
      routes.clear();
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
