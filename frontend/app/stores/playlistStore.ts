import type { Playlist } from "~/types/playlist";

export const usePlaylistStore = defineStore("playlistStore", () => {
  const {
    data: playlist,
    pending: loading,
    error,
    refresh: fetch,
  } = useAPI<Playlist[]>("/playlist/get-all");

  const route = useRoute();
  const playlistMap = ref<Map<string, Playlist>>();
  const currentPlaylistName = computed<string>(() => {
    const list = route.query["list"];

    if (!list) return "All";

    return String(list);
  });

  watch([playlist, loading], ([playlist, loading]) => {
    if (loading || !playlist) return;

    playlistMap.value = new Map(playlist.map((p) => [p.name, p]));
  });

  function isMusicInPlaylist(uuid: string, playlistName: string) {
    const playlist = playlistMap.value?.get(playlistName);

    if (!playlist) return null;

    const index = playlist.playlist_items.findIndex((p) => p.music_id == uuid);

    if (index == -1) return null;

    return playlist.playlist_items[index];
  }

  return {
    playlist,
    playlistMap,
    currentPlaylistName,
    loading,
    error,
    fetch,
    isMusicInPlaylist,
  };
});
