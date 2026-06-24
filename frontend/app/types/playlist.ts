import type { Music } from "./music";

export interface Playlist {
  id: number;
  name: string;
  created_at: string;
  playlist_items: PlaylistItem[];
}

export interface PlaylistItem {
  id: number;
  playlist_id: number;
  music_id: string;
  order_pos: string;
  music: Music;
}
