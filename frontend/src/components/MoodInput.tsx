import React, { useState } from 'react';
import { MoodType } from '../types';
import { moodOptions } from '../data/moodTypes';

interface MoodInputProps {
    onSubmit: (mood: MoodType, note?: string) => void;
}

const MoodInput: React.FC<MoodInputProps> = ({ onSubmit }) => {
    const [selectedMood, setSelectedMood] = useState<MoodType | null>(null);
    const [note, setNote] = useState('');

    const handleSubmit = () => {
        if (!selectedMood) return;
        onSubmit(selectedMood, note);
        setSelectedMood(null);
        setNote('');
    };

    return (
        <div className="p-4 rounded-xl border border-gray-300 dark:border-gray-700 bg-gray-50 dark:bg-gray-800 shadow-sm space-y-4">
            <h2 className="text-xl font-semibold">How are you feeling today?</h2>

            <div className="flex flex-wrap gap-3">
                {Object.entries(moodOptions).map(([key, { emoji, label }]) => {
                    const moodKey = key as MoodType;
                    const isSelected = selectedMood === moodKey;

                    return (
                        <button
                            key={moodKey}
                            type="button"
                            onClick={() => setSelectedMood(moodKey)}
                            className={`text-2xl p-2 rounded-full border ${
                                isSelected
                                    ? 'bg-blue-100 border-blue-400 dark:bg-blue-800 dark:border-blue-300'
                                    : 'border-gray-300 dark:border-gray-600 hover:bg-gray-200 dark:hover:bg-gray-700'
                            }`}
                            title={label}
                            aria-label={label}
                        >
                            {emoji}
                        </button>
                    );
                })}
            </div>

            <textarea
                value={note}
                onChange={(e) => setNote(e.target.value)}
                placeholder="Anything you'd like to note?"
                className="w-full p-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-900"
                rows={3}
            />

            <button
                onClick={handleSubmit}
                disabled={!selectedMood}
                className={`w-full py-2 px-4 rounded-lg font-semibold transition ${
                    selectedMood
                        ? 'bg-blue-600 text-white hover:bg-blue-700'
                        : 'bg-gray-300 text-gray-600 cursor-not-allowed'
                }`}
            >
                Log Mood
            </button>
        </div>
    );
};

export default MoodInput;
