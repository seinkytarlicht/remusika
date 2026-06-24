import type { Music } from "~/types/music";

export const usePlayerStore = defineStore("playerStore", () => {
  const route = useRoute();
  const musicStore = useMusicStore();

  const JUMP_AUDIO = 5; // seconds
  const isPlaying = ref<boolean>(false);
  const audioRef = ref<HTMLAudioElement>();
  const nowPlaying = ref<Music>();
  const duration = ref(0);
  const currentTime = ref(0);
  const isSeeking = ref(false);
  const isLooping = ref(false);
  const isEnded = ref(false);

  watch([audioRef], ([audio]) => {
    if (!audio) return;

    audio.addEventListener("loadedmetadata", () => {
      duration.value = audio.duration;
    });

    audio.addEventListener("ended", () => {
      isEnded.value = true;
    });

    audio.addEventListener("timeupdate", () => {
      if (!isSeeking.value) currentTime.value = audio.currentTime;
    });
  });

  function prevMusic() {
    musicStore.getPrevSong();
    isEnded.value = false;
    if (nowPlaying.value?.uuid == musicStore.currentMusic?.uuid) {
      audioRef.value?.play();
    } else {
      navigateTo({
        name: "play",
        query: {
          ...route.query,
          m: musicStore.currentMusic?.uuid,
        },
      });
    }
  }

  function nextMusic() {
    musicStore.getNextSong();
    isEnded.value = false;
    if (nowPlaying.value?.uuid == musicStore.currentMusic?.uuid) {
      audioRef.value?.play();
    } else {
      navigateTo({
        name: "play",
        query: {
          ...route.query,
          m: musicStore.currentMusic?.uuid,
        },
      });
    }
  }

  function playMusic() {
    if (!audioRef.value) return;

    if (audioRef.value.paused) {
      audioRef.value.play();
    } else {
      audioRef.value.pause();
    }
  }

  function rewind() {
    if (!audioRef.value) return;

    audioRef.value.currentTime -= JUMP_AUDIO;
  }

  function fastForward() {
    if (!audioRef.value) return;

    audioRef.value.currentTime += JUMP_AUDIO;
  }

  function startSeek() {
    if (!audioRef.value) return;

    isSeeking.value = true;

    audioRef.value.pause();
  }

  function seekMusic(value: number | undefined) {
    if (!value) return;

    currentTime.value = Number(value);

    if (audioRef.value) {
      audioRef.value.currentTime = Number(value);
    }
  }

  function endSeek() {
    isSeeking.value = false;

    if (audioRef.value?.paused) {
      audioRef.value?.play();
    }
  }

  function toggleLoop() {
    if (audioRef.value) {
      audioRef.value.loop = !audioRef.value.loop;
      isLooping.value = audioRef.value.loop;
    }
  }

  return {
    nowPlaying,
    isPlaying,
    audioRef,
    currentTime,
    duration,
    isLooping,
    isEnded,
    nextMusic,
    prevMusic,
    playMusic,
    fastForward,
    rewind,
    startSeek,
    seekMusic,
    endSeek,
    toggleLoop,
  };
});
