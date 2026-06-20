import { defineStore } from "pinia";
import type { Music } from "~/types/music";

export const useMusicStore = defineStore("musicStore", () => {
  const {
    data: music,
    pending: loading,
    error,
    refresh: fetch,
  } = useAPI<Music[]>("/music/get-all");

  const playlistStore = usePlaylistStore();

  const currentMusic = ref<Music>();
  const currentMusicError = ref(false);

  const searchQueue = ref<string>("");
  const searchQueueDebounce = useDebounce(searchQueue, 300);
  const searchShowed = ref<string>("");
  const searchShowedDebounce = useDebounce(searchShowed, 300);

  const showedPlaylist = computed<Music[]>(() => {
    if (!music.value) return [];

    const showedPlaylist = playlistStore.showedPlaylist;
    const showedPlaylistName = playlistStore.showedPlaylistName;

    let listSong = music.value;
    if (showedPlaylistName != "All") {
      listSong = showedPlaylist;
    }

    const search = searchShowedDebounce.value;

    if (search) {
      listSong = listSong.filter((m) => {
        return (
          m.title.toLowerCase().includes(search.toLowerCase()) ||
          m.artist.toLowerCase().includes(search.toLowerCase()) ||
          m.album.toLowerCase().includes(search.toLowerCase())
        );
      });
    }

    return listSong;
  });

  const queue = computed<Music[]>(() => {
    if (!music.value) return [];

    const selectedPlaylist = playlistStore.selectedPlaylist;
    const selectedPlaylistName = playlistStore.selectedPlaylistName;

    let listSong = music.value;
    if (selectedPlaylistName != "All") {
      listSong = selectedPlaylist;
    }

    return listSong;
  });

  const queueView = computed<Music[]>(() => {
    let listSong = queue.value;

    const search = searchQueueDebounce.value;

    if (search) {
      listSong = listSong.filter((m) => {
        return (
          m.title.toLowerCase().includes(search.toLowerCase()) ||
          m.artist.toLowerCase().includes(search.toLowerCase()) ||
          m.album.toLowerCase().includes(search.toLowerCase())
        );
      });
    }

    return listSong;
  });

  function setCurrentMusic(uuid: string) {
    if (loading.value || !music.value) {
      currentMusicError.value = true;
      currentMusic.value = undefined;
      return undefined;
    }

    let curMusic = music.value.find((item) => item.uuid == uuid);

    if (!curMusic) {
      currentMusicError.value = true;
      currentMusic.value = undefined;
      return undefined;
    }

    currentMusicError.value = false;
    currentMusic.value = curMusic;
    return curMusic;
  }

  function getPrevSong() {
    if (loading.value && !queue.value) return;

    const currentIndex = queue.value!.findIndex(
      (item) => item.uuid === currentMusic.value?.uuid,
    );

    if (currentIndex === -1) return null;

    if (currentIndex === 0) {
      currentMusic.value = queue.value![queue.value!.length - 1];
    } else {
      currentMusic.value = queue.value![currentIndex - 1];
    }

    return currentMusic.value;
  }

  function getNextSong() {
    if (loading.value && !queue.value) return;

    const currentIndex = queue.value!.findIndex(
      (item) => item.uuid === currentMusic.value?.uuid,
    );

    if (currentIndex === -1) return null;

    if (currentIndex === queue.value!.length - 1) {
      currentMusic.value = queue.value![0];
    } else {
      currentMusic.value = queue.value![currentIndex + 1];
    }

    return currentMusic.value;
  }

  return {
    currentMusic,
    currentMusicError,
    music: queueView,
    showedPlaylist,
    searchQueue,
    searchShowed,
    loading,
    error,
    fetch,
    setCurrentMusic,
    getPrevSong,
    getNextSong,
  };
});
