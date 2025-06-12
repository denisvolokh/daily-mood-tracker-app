import React from 'react';
import { MoodEntry, MoodType } from '../types';
import { moodOptions } from '../data/moodTypes';

interface MoodStatsProps {
    entries: MoodEntry[];
}

const MoodStats: React.FC<MoodStatsProps> = ({ entries }) => {
    if (entries.length === 0) return null;

    // Count occurrences
    const moodCounts: Record<MoodType, number> = {
        happy: 0,
        neutral: 0,
        sad: 0,
        angry: 0,
        tired: 0,
    };

    entries.forEach((entry) => {
        moodCounts[entry.mood]++;
    });

    const total = entries.length;

    // Very naive "average" mood rating
    const moodScore: Record<MoodType, number> = {
        happy: 3,
        neutral: 2,
        tired: 1.5,
        sad: 1,
        angry: 0.5,
    };

    const avgScore =
        entries.reduce((acc, entry) => acc + moodScore[entry.mood], 0) / total;

    const avgLabel =
        avgScore >= 2.5
            ? '🙂 Mostly Positive'
            : avgScore >= 1.5
                ? '😐 Mixed'
                : '☁️ Mostly Negative';

    return (
        <div className="rounded-xl border border-gray-300 dark:border-gray-700 bg-gray-50 dark:bg-gray-800 shadow-sm p-4 space-y-4">
            <h2 className="text-lg font-semibold">Mood Stats</h2>

            <div className="flex flex-wrap gap-3 text-lg">
                {Object.entries(moodCounts).map(([key, count]) =>
                    count > 0 ? (
                        <div key={key} className="flex items-center gap-2">
                            <span>{moodOptions[key as MoodType].emoji}</span>
                            <span>× {count}</span>
                        </div>
                    ) : null
                )}
            </div>

            <div className="text-sm text-gray-700 dark:text-gray-300">
                <strong>Average Mood:</strong> {avgLabel}
            </div>
        </div>
    );
};

export default MoodStats;
