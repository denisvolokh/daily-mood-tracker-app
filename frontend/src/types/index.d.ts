export type MoodType = "happy" | "neutral" | "sad" | "angry" | "tired";

export interface MoodEntry {
    date: string; // ISO string
    mood: MoodType;
    note?: string;
}
