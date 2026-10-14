import { useCallback, useEffect, useRef, useState } from "react";
import type {
    MatchingWords,
    MatchFeedback,
    MatchingGameState,
} from "./matching.words.type";
import "./MatchingWordsPage.css";
import { BankWordService } from "../../features/bank_word/services/BankWordService";

const MAX_PAIRS = 6;
const GAME_DURATION_SECONDS = 60;

const WRONG_FEEDBACK_DURATION = 250;
const CORRECT_FEEDBACK_DURATION = 450;

/**
 * Mengacak urutan array tanpa mengubah array aslinya.
 * Menggunakan algoritma Fisher-Yates.
 */
const shuffleArray = <T,>(items: T[]): T[] => {
    const shuffled = [...items];

    for (let i = shuffled.length - 1; i > 0; i--) {
        const randomIndex = Math.floor(
            Math.random() * (i + 1)
        );

        [shuffled[i], shuffled[randomIndex]] = [
            shuffled[randomIndex],
            shuffled[i],
        ];
    }

    return shuffled;
};

/**
 * Membuat state permainan dari data API.
 * Maksimal enam pasangan ditampilkan dalam satu grup.
 */
const createGameState = (
    data: MatchingWords[]
): MatchingGameState => {
    const allWords = shuffleArray(data);
    const firstGroup = allWords.slice(0, MAX_PAIRS);
    const queue = allWords.slice(MAX_PAIRS);

    return {
        words: firstGroup,
        rightWords: shuffleArray(firstGroup),
        queue,
        totalPairs: allWords.length,
    };
};

/**
 * Membuat state permainan kosong.
 */
const createEmptyGameState = (): MatchingGameState => ({
    words: [],
    rightWords: [],
    queue: [],
    totalPairs: 0,
});

export const MatchingWordsPage = () => {
    const [game, setGame] = useState<MatchingGameState>(createEmptyGameState);

    const [isReady, setIsReady] = useState(false);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const [selectedLeft, setSelectedLeft] = useState<string | null>(null);
    const [selectedRight, setSelectedRight] = useState<string | null>(null);
    const [hintEnabled, setHintEnabled] = useState(false);
    const [timeLeft, setTimeLeft] = useState(GAME_DURATION_SECONDS);
    const [feedback, setFeedback] = useState<MatchFeedback | null>(null);
    const feedbackTimer = useRef<ReturnType<typeof setTimeout> | null>(null);

    const {
        words,
        rightWords,
        queue,
        totalPairs,
    } = game;

    const isGameComplete =
        isReady &&
        words.length === 0 &&
        queue.length === 0;

    const isTimeUp = isReady && timeLeft === 0;
    const isGameOver = isGameComplete || isTimeUp;
    const matchedPairs = totalPairs - words.length - queue.length;

    /**
     * Memuat permainan baru ketika pengguna memilih Try Again
     * atau Play Again.
     */
    const loadGame = useCallback(async () => {
        setLoading(true);
        setError(null);
        setIsReady(false);

        if (feedbackTimer.current !== null) {
            clearTimeout(feedbackTimer.current);
            feedbackTimer.current = null;
        }

        setGame(createEmptyGameState());
        setSelectedLeft(null);
        setSelectedRight(null);
        setFeedback(null);
        setHintEnabled(false);
        setTimeLeft(GAME_DURATION_SECONDS);

        try {
            const matchingWords =
                await BankWordService.getMatchingWordsContent();

            if (matchingWords.length === 0) {
                setError(
                    "No matching words available. Please try again later."
                );
                return;
            }

            setGame(createGameState(matchingWords));
            setIsReady(true);
        } catch (error) {
            console.error("Failed to load matching words:", error);
            setError("Failed to load matching words. Please try again.");
        } finally {
            setLoading(false);
        }
    }, []);

    /**
     * WAJIB: memuat data secara otomatis saat halaman dibuka.
     * Tidak memerlukan klik dari pengguna.
     */
    useEffect(() => {
        let cancelled = false;

        const fetchInitialGame = async () => {
            try {
                // Langsung meminta data API.
                // loading sudah true sejak inisialisasi state.
                const matchingWords = await BankWordService.getMatchingWordsContent();

                // Abaikan respons jika effect sudah dibersihkan.
                if (cancelled) return;

                if (matchingWords.length === 0) {
                    setError("No matching words available. Please try again later.");
                    return;
                }

                // Data berhasil diterima: langsung mulai permainan.
                setGame(createGameState(matchingWords));
                setIsReady(true);
            } catch (error) {
                if (cancelled) {
                    return;
                }

                console.error(
                    "Failed to load matching words on initial render:",
                    error
                );

                setError(
                    "Failed to load matching words. Please try again."
                );
            } finally {
                if (!cancelled) {
                    setLoading(false);
                }
            }
        };

        fetchInitialGame();

        return () => {
            cancelled = true;
        };
    }, []);

    /**
     * Timer berjalan setelah permainan siap,
     * dan berhenti ketika permainan selesai atau waktu habis.
     */
    useEffect(() => {
        if (!isReady || isGameOver) {
            return;
        }

        const timer = setInterval(() => {
            setTimeLeft((current) =>
                Math.max(0, current - 1)
            );
        }, 1000);

        return () => clearInterval(timer);
    }, [isReady, isGameOver]);

    /**
     * Membersihkan feedback timeout ketika komponen dilepas.
     */
    useEffect(() => {
        return () => {
            if (feedbackTimer.current !== null) {
                clearTimeout(feedbackTimer.current);
            }
        };
    }, []);

    /**
     * Memeriksa apakah dua kata cocok.
     */
    const checkMatch = (
        leftId: string,
        rightId: string
    ) => {
        if (feedback !== null || !isReady || isGameOver) {
            return;
        }

        const isCorrect = leftId === rightId;

        let leftIds = [leftId];
        let rightIds = [rightId];

        if (!isCorrect && hintEnabled) {
            leftIds = [leftId, rightId];
            rightIds = [rightId, leftId];
        }

        setFeedback({
            type: isCorrect ? "correct" : "wrong",
            leftIds: isCorrect ? [leftId] : leftIds,
            rightIds: isCorrect ? [rightId] : rightIds,
        });

        if (feedbackTimer.current !== null) {
            clearTimeout(feedbackTimer.current);
        }

        const duration = isCorrect
            ? CORRECT_FEEDBACK_DURATION
            : WRONG_FEEDBACK_DURATION;

        feedbackTimer.current = setTimeout(() => {
            if (isCorrect) {
                setGame((current) => {
                    const remainingWords =
                        current.words.filter(
                            (word) => word.word_id !== leftId
                        );

                    const remainingRightWords =
                        current.rightWords.filter(
                            (word) => word.word_id !== rightId
                        );

                    // Jika grup saat ini selesai, tampilkan grup berikutnya.
                    if (
                        remainingWords.length === 0 &&
                        current.queue.length > 0
                    ) {
                        const nextGroup =
                            current.queue.slice(0, MAX_PAIRS);

                        return {
                            ...current,
                            words: nextGroup,
                            rightWords: shuffleArray(nextGroup),
                            queue: current.queue.slice(MAX_PAIRS),
                        };
                    }

                    return {
                        ...current,
                        words: remainingWords,
                        rightWords: remainingRightWords,
                    };
                });
            }

            setSelectedLeft(null);
            setSelectedRight(null);
            setFeedback(null);
            feedbackTimer.current = null;
        }, duration);
    };

    /**
     * Menangani pilihan kata bahasa Inggris.
     */
    const handleLeftClick = (wordId: string) => {
        if (feedback !== null || !isReady || isGameOver) {
            return;
        }

        if (selectedLeft === wordId) {
            setSelectedLeft(null);
            return;
        }

        setSelectedLeft(wordId);

        if (selectedRight !== null) {
            checkMatch(wordId, selectedRight);
        }
    };

    /**
     * Menangani pilihan terjemahan.
     */
    const handleRightClick = (wordId: string) => {
        if (feedback !== null || !isReady || isGameOver) {
            return;
        }

        if (selectedRight === wordId) {
            setSelectedRight(null);
            return;
        }

        setSelectedRight(wordId);

        if (selectedLeft !== null) {
            checkMatch(selectedLeft, wordId);
        }
    };

    /**
     * Menentukan class visual tombol.
     */
    const getButtonClassName = (
        wordId: string,
        side: "left" | "right"
    ) => {
        const selectedId =
            side === "left" ? selectedLeft : selectedRight;

        const feedbackIds =
            side === "left"
                ? feedback?.leftIds
                : feedback?.rightIds;

        const isSelected = selectedId === wordId;

        const isCorrect =
            feedback?.type === "correct" &&
            feedbackIds?.includes(wordId);

        const isWrong =
            feedback?.type === "wrong" &&
            feedbackIds?.includes(wordId);

        return [
            "matching-button",
            isSelected ? "is-selected" : "",
            isCorrect ? "is-correct" : "",
            isWrong ? "is-wrong" : "",
        ]
            .filter(Boolean)
            .join(" ");
    };

    const formattedTime = `${Math.floor(timeLeft / 60)
        .toString()
        .padStart(2, "0")}:${(timeLeft % 60)
        .toString()
        .padStart(2, "0")}`;

    /**
     * Tampilan loading ketika data sedang dimuat.
     */
    if (loading) {
        return (
            <main className="matching-page">
                <header className="matching-header">
                    <h1>Matching Words</h1>
                    <p>
                        Match each English word with its correct translation.
                    </p>
                </header>

                <section
                    className="matching-loading"
                    role="status"
                    aria-live="polite"
                >
                    <span
                        className="matching-spinner"
                        aria-hidden="true"
                    />
                    <p>Loading matching words...</p>
                </section>
            </main>
        );
    }

    /**
     * Tampilan error jika API gagal atau tidak ada data.
     */
    if (error || !isReady) {
        return (
            <main className="matching-page">
                <header className="matching-header">
                    <h1>Matching Words</h1>
                    <p>
                        Match each English word with its correct translation.
                    </p>
                </header>

                <section className="matching-loading">
                    <p role="alert">
                        {error ?? "No matching words available."}
                    </p>

                    <button
                        type="button"
                        onClick={() => void loadGame()}
                        disabled={loading}
                    >
                        Try Again
                    </button>
                </section>
            </main>
        );
    }

    return (
        <main className="matching-page">
            <header className="matching-header">
                <h1>Matching Words</h1>
                <p>
                    Match each English word with its correct translation.
                </p>
            </header>

            <section
                className="matching-toolbar"
                aria-label="Game settings and progress"
            >
                <label className="matching-hint-control">
                    <input
                        type="checkbox"
                        checked={hintEnabled}
                        onChange={(event) =>
                            setHintEnabled(event.target.checked)
                        }
                        disabled={isGameOver}
                    />

                    <span>Enable Hint</span>
                </label>

                <div
                    className={`matching-timer ${
                        timeLeft <= 10 && !isGameOver
                            ? "is-urgent"
                            : ""
                    }`}
                    role="timer"
                    aria-label={`Time remaining: ${formattedTime}`}
                >
                    <span className="matching-timer-label">
                        Time
                    </span>

                    <span className="matching-timer-value">
                        {formattedTime}
                    </span>
                </div>

                <p className="matching-progress">
                    Matched: {matchedPairs} / {totalPairs}
                </p>
            </section>

            <section
                className="matching-game"
                aria-label="Word matching game"
            >
                <div className="matching-column">
                    <h2>Words</h2>

                    <div className="matching-list">
                        {words.map((word) => (
                            <button
                                key={word.word_id}
                                type="button"
                                className={getButtonClassName(
                                    word.word_id,
                                    "left"
                                )}
                                aria-pressed={
                                    selectedLeft === word.word_id
                                }
                                disabled={
                                    feedback !== null || isGameOver
                                }
                                onClick={() =>
                                    handleLeftClick(word.word_id)
                                }
                            >
                                {word.word}
                            </button>
                        ))}
                    </div>
                </div>

                <div className="matching-column">
                    <h2>Translations</h2>

                    <div className="matching-list">
                        {rightWords.map((word) => (
                            <button
                                key={word.word_id}
                                type="button"
                                className={getButtonClassName(
                                    word.word_id,
                                    "right"
                                )}
                                aria-pressed={
                                    selectedRight === word.word_id
                                }
                                disabled={
                                    feedback !== null || isGameOver
                                }
                                onClick={() =>
                                    handleRightClick(word.word_id)
                                }
                            >
                                {word.translations[0] || "-"}
                            </button>
                        ))}
                    </div>
                </div>
            </section>

            <footer className="matching-footer">
                {isGameComplete ? (
                    <>
                        <p
                            className="matching-status matching-status-success"
                            role="status"
                        >
                            Excellent! All pairs matched.
                        </p>

                        <button
                            type="button"
                            onClick={() => void loadGame()}
                            disabled={loading}
                        >
                            Play Again
                        </button>
                    </>
                ) : isTimeUp ? (
                    <>
                        <p
                            className="matching-status matching-status-timeout"
                            role="status"
                        >
                            Time's up! You matched{" "}
                            {matchedPairs} of {totalPairs} pairs.
                        </p>

                        <button
                            type="button"
                            onClick={() => void loadGame()}
                            disabled={loading}
                        >
                            Try Again
                        </button>
                    </>
                ) : (
                    <p className="matching-status">
                        Match all the words before time runs out!
                    </p>
                )}
            </footer>
        </main>
    );
};