import type { Music } from "~/types/music";

export function useMediaSession() {
  const initSession = (music: Music) => {
    if (!("mediaSession" in navigator)) return;

    navigator.mediaSession.metadata = new MediaMetadata({
      title: music.title,
      artist: music.artist,
      album: music.album,
      artwork: [
        {
          src: music.image_url,
          sizes: "512x512",
          type: "image/jpeg",
        },
      ],
    });

    // // 3. Optional: Define action handlers (e.g., play, pause, seek)
    // navigator.mediaSession.setActionHandler("play", () => {
    //   // Trigger your Vue/Audio play logic
    // });

    // navigator.mediaSession.setActionHandler("pause", () => {
    //   // Trigger your Vue/Audio pause logic
    // });
  };

  // const updatePlaybackState = (state) => {
  //   if ("mediaSession" in navigator) {
  //     navigator.mediaSession.playbackState = state; // 'none' | 'paused' | 'playing'
  //   }
  // };

  return { initSession };
}
