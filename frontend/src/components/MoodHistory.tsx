import {MoodEntry} from "../types";
import {moodOptions} from "../data/moodTypes";

interface MoodHistoryProps {
    entries: MoodEntry[];
}

export const MoodHistory: React.FC<MoodHistoryProps> = ({entries}) => {
    if (entries.length === 0) {
        return (
            <div className="text-center text-gray-500 dark:text-gray-400">
                No mood entries yet.
            </div>
        );
    }

    return (
        <div className="rounded-xl border border-gray-300 dark:border-gray-700 bg-gray-50 dark:bg-gray-800 shadow-sm">
            <h2 className="text-lg font-semibold px-4 py-3 border-b border-gray-200 dark:border-gray-600">
                Mood History
            </h2>
            <ul className="divide-y divide-gray-200 dark:divide-gray-700 max-h-64 overflow-y-auto">
                {entries.map((entry, index) => {
                    const mood = moodOptions[entry.mood];
                    const date = new Date(entry.date).toLocaleDateString(undefined, {
                        year: 'numeric',
                        month: 'short',
                        day: 'numeric',
                    });

                    return (
                        <li
                            key={index}
                            className="px-4 py-3 flex justify-between items-start gap-4 bg-white dark:bg-gray-900 even:bg-gray-50 dark:even:bg-gray-800"
                        >
                            <div>
                                <div className="font-medium">{date}</div>
                                {entry.note && (
                                    <div className="text-sm text-gray-600 dark:text-gray-400 mt-1">
                                        “{entry.note}”
                                    </div>
                                )}
                            </div>
                            <div className="text-2xl" title={mood.label}>
                                {mood.emoji}
                            </div>
                        </li>
                    );
                })}
            </ul>
        </div>
    )
}

export default MoodHistory;