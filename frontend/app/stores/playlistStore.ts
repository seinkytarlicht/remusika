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

  const playlistMap = computed<Map<number, Playlist> | undefined>(() => {
    if (loading.value || !playlist.value) return;

    return new Map(playlist.value.map((p) => [p.id, p]));
  });

  const selectedPlaylistName = computed<string>(() => {
    if (!playlistMap.value) return "";
    if (showedPlaylistId.value == 0) return "All";

    const playlist = playlistMap.value.get(selectedPlaylistId.value);
    if (!playlist) return "All";

    return playlist.name;
  });
  const selectedPlaylistId = computed<number>(() => {
    const showlist = route.query["playlist"];

    if (!showlist) return 0;

    return Number(showlist);
  });
  const selectedPlaylist = computed<Music[]>(() => {
    if (!selectedPlaylistId.value || !playlistMap.value) return [];
    if (selectedPlaylistId.value == 0) return [];

    const playlist = playlistMap.value.get(selectedPlaylistId.value);
    if (!playlist) return [];

    return playlist?.playlist_items.map((pi) => {
      const plm = pi.music;
      plm.uuid = pi.music_id;

      return plm;
    });
  });

  const showedPlaylistName = computed<string>(() => {
    if (!playlistMap.value) return "";
    if (showedPlaylistId.value == 0) return "All";

    const playlist = playlistMap.value.get(showedPlaylistId.value);
    if (!playlist) return "All";

    return playlist.name;
  });
  const showedPlaylistId = computed<number>(() => {
    const showlist = route.query["showlist"];

    if (!showlist) return 0;

    return Number(showlist);
  });
  const showedPlaylist = computed<Music[]>(() => {
    if (!showedPlaylistId.value || !playlistMap.value) return [];
    if (showedPlaylistId.value == 0) return [];

    const playlist = playlistMap.value.get(showedPlaylistId.value);
    if (!playlist) return [];

    return playlist?.playlist_items.map((pi) => {
      const plm = pi.music;
      plm.uuid = pi.music_id;

      return plm;
    });
  });

  function isMusicInPlaylist(uuid: string, playlistId: number) {
    const playlist = playlistMap.value?.get(playlistId);

    if (!playlist) return null;

    const index = playlist.playlist_items.findIndex((p) => p.music_id == uuid);

    if (index == -1) return null;

    return playlist.playlist_items[index];
  }

  return {
    playlist,
    playlistMap,
    showedPlaylistName,
    showedPlaylistId,
    showedPlaylist,
    selectedPlaylistName,
    selectedPlaylistId,
    selectedPlaylist,
    loading,
    error,
    fetch,
    isMusicInPlaylist,
  };
});
