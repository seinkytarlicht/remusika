import { defineStore } from "pinia";
import type { Music } from "~/types/music";

export const useMusicStore = defineStore("musicStore", () => {
  const {
    data: music,
    pending: loading,
    error,
    refresh: fetch,
  } = useAPI<Music[]>("/music/get-all");

  const currentMusic = ref<Music>();
  const currentMusicError = ref(false);
  const search = ref<string>("");
  const searchDebounce = useDebounce(search, 300);
  const queue = computed<Music[]>(() => {
    if (!music.value) return [];

    let listSong = music.value;

    if (searchDebounce.value) {
      listSong = listSong.filter((m) => {
        return (
          m.title.toLowerCase().includes(searchDebounce.value.toLowerCase()) ||
          m.artist.toLowerCase().includes(searchDebounce.value.toLowerCase()) ||
          m.album.toLowerCase().includes(searchDebounce.value.toLowerCase())
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
    music: queue,
    loading,
    error,
    search,
    fetch,
    setCurrentMusic,
    getPrevSong,
    getNextSong,
  };
});
