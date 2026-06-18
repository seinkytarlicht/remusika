import type { Music } from "./music";

export interface Playlist {
  id: number;
  name: string;
  created_at: string;
  playlist_items: PlaylistItem[];
}

export interface PlaylistItem {
  id: 10;
  playlist_id: 1;
  music_id: "cbee416a3b01";
  order_pos: "1";
  music: Music;
}
