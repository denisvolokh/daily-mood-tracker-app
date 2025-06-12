import {MoodEntry, MoodType} from "./types";
import {useEffect, useState} from "react";
import MoodInput from "./components/MoodInput";
import MoodHistory from "./components/MoodHistory";
import MoodStats from "./components/MoodStats";

const LOCAL_STORAGE_KEY = "mood-tracker-app";

const App: React.FC = () => {
    const [entries, setEntries] = useState<MoodEntry[]>([]);

    // Load from Local Storage
    useEffect(() => {
        const stored = localStorage.getItem(LOCAL_STORAGE_KEY);
        if (stored) {
            try {
                setEntries(JSON.parse(stored));
            } catch (e) {
                console.error('[!!!] Failed to parse stored entries.');
            }
        }
    }, [])

    // Save to local storage
    useEffect(() => {
        localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(entries));
    }, [entries]);

    // Handle new mood submission
    const handleLogMood = (mood : MoodType, note?: string) => {
        const newEntry: MoodEntry = {
            date: new Date().toISOString(),
            mood,
            note: note?.trim() || undefined
        }
        setEntries([newEntry, ...entries]);
    }

    return (
        <main className="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-gray-100 font-sans">
            <div className="max-w-xl mx-auto px-4 py-8 space-y-10">
                <section>
                    <MoodInput onSubmit={handleLogMood} />
                </section>

                <section>
                    <MoodHistory entries={entries} />
                </section>

                <section>
                    <MoodStats entries={entries} />
                </section>
            </div>
        </main>
    );
}

export default App;