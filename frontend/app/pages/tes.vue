<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";

// Reactive state for tracking playback
const isPlaying = ref(false);
const audioRef = ref<HTMLAudioElement | null>(null);

// Initialize and manage the Media Session API
const updateMediaSession = () => {
  if (!("mediaSession" in navigator)) return;

  // 1. Set Track Metadata for Lock Screen / Notification Tray
  navigator.mediaSession.metadata = new MediaMetadata({
    title: "Epic Cinematic Track",
    artist: "Creative Producer",
    album: "Nuxt Audio Experiments",
    artwork: [
      { src: "https://placehold.co", sizes: "96x96", type: "image/png" },
      { src: "https://placehold.co", sizes: "512x512", type: "image/png" },
    ],
  });

  // 2. Set Action Handlers for System Controls
  navigator.mediaSession.setActionHandler("play", () => {
    playAudio();
  });

  navigator.mediaSession.setActionHandler("pause", () => {
    pauseAudio();
  });

  navigator.mediaSession.setActionHandler("stop", () => {
    stopAudio();
  });
};

// Playback control functions
const playAudio = async () => {
  if (!audioRef.value) return;
  try {
    await audioRef.value.play();
    isPlaying.value = true;
    if ("mediaSession" in navigator) {
      navigator.mediaSession.playbackState = "playing";
    }
  } catch (error) {
    console.error("Audio playback failed:", error);
  }
};

const pauseAudio = () => {
  if (!audioRef.value) return;
  audioRef.value.pause();
  isPlaying.value = false;
  if ("mediaSession" in navigator) {
    navigator.mediaSession.playbackState = "paused";
  }
};

const stopAudio = () => {
  if (!audioRef.value) return;
  audioRef.value.pause();
  audioRef.value.currentTime = 0;
  isPlaying.value = false;
  if ("mediaSession" in navigator) {
    navigator.mediaSession.playbackState = "none";
  }
};

// Safely invoke client-side APIs during browser execution
onMounted(() => {
  updateMediaSession();
});

// Clean up action handlers when the component is unmounted
onBeforeUnmount(() => {
  if ("mediaSession" in navigator) {
    navigator.mediaSession.setActionHandler("play", null);
    navigator.mediaSession.setActionHandler("pause", null);
    navigator.mediaSession.setActionHandler("stop", null);
  }
});
</script>

<template>
  <div class="audio-player">
    <h1>Nuxt Media Session Example</h1>

    <!-- Standard HTML5 Audio Element -->
    <audio
      ref="audioRef"
      src="https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3"
      @play="isPlaying = true"
      @pause="isPlaying = false"
    ></audio>

    <div class="controls">
      <button v-if="!isPlaying" @click="playAudio" class="btn-play">
        ▶ Play Track
      </button>
      <button v-else @click="pauseAudio" class="btn-pause">
        ⏸ Pause Track
      </button>
      <button @click="stopAudio" class="btn-stop">⏹ Stop</button>
    </div>
  </div>
</template>

<style scoped>
.audio-player {
  max-width: 400px;
  margin: 2rem auto;
  padding: 2rem;
  border: 1px solid #00dc82;
  border-radius: 8px;
  text-align: center;
}
.controls {
  margin-top: 1.5rem;
  display: flex;
  gap: 10px;
  justify-content: center;
}
button {
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
}
</style>
