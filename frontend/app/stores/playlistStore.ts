import type { Music } from "~/types/music";
import type { Playlist } from "~/types/playlist";

export const usePlaylistStore = defineStore("playlistStore", () => {
  const {
    data: playlist,
    pending: loading,
    error,
    refresh: fetch,
  } = useAPI<Playlist[]>("/playlist/get-all");

  const route = useRoute();

  const playlistMap = computed<Map<string, Playlist> | undefined>(() => {
    if (loading.value || !playlist.value) return;

    return new Map(playlist.value.map((p) => [p.name, p]));
  });

  const selectedPlaylistName = computed<string>(() => {
    const showlist = route.query["playlist"];

    if (!showlist) return "All";

    return String(showlist);
  });
  const selectedPlaylist = computed<Music[]>(() => {
    if (!selectedPlaylistName.value || !playlistMap.value) return [];
    if (selectedPlaylistName.value == "All") return [];

    const playlist = playlistMap.value.get(selectedPlaylistName.value);
    if (!playlist) return [];

    return playlist?.playlist_items.map((pi) => {
      const plm = pi.music;
      plm.uuid = pi.music_id;

      return plm;
    });
  });

  const showedPlaylistName = computed<string>(() => {
    const showlist = route.query["showlist"];

    if (!showlist) return "All";

    return String(showlist);
  });
  const showedPlaylist = computed<Music[]>(() => {
    if (!showedPlaylistName.value || !playlistMap.value) return [];
    if (showedPlaylistName.value == "All") return [];

    const playlist = playlistMap.value.get(showedPlaylistName.value);
    if (!playlist) return [];

    return playlist?.playlist_items.map((pi) => {
      const plm = pi.music;
      plm.uuid = pi.music_id;

      return plm;
    });
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
    showedPlaylistName,
    showedPlaylist,
    selectedPlaylistName,
    selectedPlaylist,
    loading,
    error,
    fetch,
    isMusicInPlaylist,
  };
});
