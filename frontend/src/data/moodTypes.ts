import { MoodType } from "../types";

export const moodOptions: Record<MoodType, { label: string; emoji: string }> = {
    happy: { label: "Happy", emoji: "😀" },
    neutral: { label: "Neutral", emoji: "😐" },
    sad: { label: "Sad", emoji: "😢" },
    angry: { label: "Angry", emoji: "😠" },
    tired: { label: "Tired", emoji: "😴" },
};
