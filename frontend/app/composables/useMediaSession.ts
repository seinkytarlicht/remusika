import type { Music } from "~/types/music";

export function useMediaSession() {
  const initSession = (music: Music) => {
    if (!("mediaSession" in navigator)) return;

    navigator.mediaSession.metadata = new MediaMetadata({
      title: music.title,
      artist: music.artist,
      album: music.album,
      artwork: music.image_url
        ? [
            {
              src: music.image_url,
              sizes: "512x512",
              type: "image/jpeg",
            },
          ]
        : [],
    });
  };

  const updatePlaybackState = (state: "none" | "paused" | "playing") => {
    if ("mediaSession" in navigator) {
      navigator.mediaSession.playbackState = state;
    }
  };
  const setPrevTrack = (callback: () => void) => {
    if ("mediaSession" in navigator) {
      navigator.mediaSession.setActionHandler("previoustrack", callback);
    }
  };
  const setNextTrack = (callback: () => void) => {
    if ("mediaSession" in navigator) {
      navigator.mediaSession.setActionHandler("nexttrack", callback);
    }
  };

  return { initSession, updatePlaybackState, setPrevTrack, setNextTrack };
}
