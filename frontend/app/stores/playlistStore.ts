import type { Playlist } from "~/types/playlist";

export const usePlaylistStore = defineStore("playlistStore", () => {
  const {
    data: playlist,
    pending: loading,
    error,
    refresh: fetch,
  } = useAPI<Playlist[]>("/playlist/get-all");

  return {
    playlist,
    loading,
    error,
    fetch,
  };
});
