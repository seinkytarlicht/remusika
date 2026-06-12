export const usePlayerStore = defineStore("playerStore", () => {
  const isPlaying = ref<boolean>(false);
  const audioRef = ref<HTMLAudioElement>();
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
      console.log("ended");
      isEnded.value = true;
    });

    audio.addEventListener("timeupdate", () => {
      if (!isSeeking.value) currentTime.value = audio.currentTime;
    });
  });

  function playMusic() {
    if (!audioRef.value) return;

    if (audioRef.value.paused) {
      audioRef.value.play();
    } else {
      audioRef.value.pause();
    }
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
    isPlaying,
    audioRef,
    currentTime,
    duration,
    isLooping,
    isEnded,
    playMusic,
    startSeek,
    seekMusic,
    endSeek,
    toggleLoop,
  };
});
