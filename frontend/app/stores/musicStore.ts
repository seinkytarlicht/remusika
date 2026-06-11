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
  const queue = ref<Music[]>();
  const search = ref<string>();

  watch([music, search], ([music, search]) => {
    if (!music) return;

    let listSong = music;

    if (search) {
      listSong = listSong.filter((m) => {
        return (
          m.title.toLowerCase().includes(search.toLowerCase()) ||
          m.artist.toLowerCase().includes(search.toLowerCase()) ||
          m.album.toLowerCase().includes(search.toLowerCase())
        );
      });
    }

    queue.value = listSong;
  });

  function setCurrentMusic(uuid: string) {
    if (!loading.value && music.value) {
      let curMusic = music.value.find((item) => item.uuid == uuid);
      currentMusic.value = curMusic;
      return curMusic;
    }

    return undefined;
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
