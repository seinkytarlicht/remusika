export default defineNuxtPlugin((nuxtApp) => {
  const api = $fetch.create({
    baseURL: "/api",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return {
    provide: {
      api,
    },
  };
});
