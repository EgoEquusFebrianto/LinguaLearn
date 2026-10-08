import { useEffect, useRef, useState } from 'react';
import { UserLayout } from '../../layouts/user/UserLayout';
import { useInfiniteDictionary } from '../../features/bank_word/hooks/useInfiniteDictionary';
import "./UserDictionary.css";

export const UserDictionary = () => {
    const [activeTab, setActiveTab] = useState("all");
    const [isSortOpen, setIsSortOpen] = useState(false);
    const [expandedCard, setExpandedCard] = useState<string | null>(null);

    const observerTarget = useRef<HTMLDivElement | null>(null);

    const {
        dictionaryContent,
        isLoading,
        isFetchingNextPage,
        hasNextPage,
        fetchNextPage,
    } = useInfiniteDictionary(); 


    const handleCardToggle = (senseId: string) => {
        setExpandedCard((current) => current === senseId ? null : senseId);
    };

    const handleCardClose = () => {
        setExpandedCard(null)
    };

    const handleSortToggle = () => {
        setIsSortOpen((current) => !current);
    };

    const handleTabChange = (tab: string) => {
        setActiveTab(tab);
    };

    useEffect(() => {
        const target = observerTarget.current;

        if (!target) return;

        const observer = new IntersectionObserver(
            (entries) => {
                const entry = entries[0]

                if (
                    entry.isIntersecting &&
                    hasNextPage &&
                    !isFetchingNextPage
                ) {
                    fetchNextPage();
                }
            },
            {
                root: null,
                rootMargin: "200px",
                threshold: 0,
            }
        );

        observer.observe(target);

        return () => {
            observer.disconnect();
        };
    }, [hasNextPage, isFetchingNextPage, fetchNextPage]);

    return (
        <UserLayout>
            <main className="dictionary-page">
                {/* Page Header */}
                <header className="dictionary-header">
                    <h1>My Dictionary</h1>
                </header>

                {/* Dictionary Tabs */}
                <nav 
                    className="dictionary-tabs"
                    aria-label="Dictionary filters"
                >
                    <button 
                        type="button" 
                        className={`tab-button ${activeTab === "all" ? "active" : ""}`}
                        onClick={() => handleTabChange("all")}
                    >
                        All
                    </button>

                    <button 
                        type="button" 
                        className={`tab-button ${activeTab === "under-study" ? "active" : ""}`}
                        onClick={() => handleTabChange("under-study")}
                    >
                        Under Study
                    </button>

                    <button 
                        type="button" 
                        className={`tab-button ${activeTab === "learned" ? "active" : ""}`}
                        onClick={() => handleTabChange("learned")}
                    >
                        Has Been Learned
                    </button>
                </nav>

                {/* Dictionary Controls */}
                <section className="dictionary-controls">
                    <button
                        type="button"
                        className="sort-button"
                        aria-expanded="false"
                        aria-controls="sort-options"
                        onClick={handleSortToggle}
                    >
                        Sort Options
                    </button>

                    <div 
                        id="sort-options"
                        className={`sort-options ${isSortOpen ? "open" : ""}`}
                    >
                        <p>Sort by</p>

                        <label>
                            <input
                                type="radio"
                                name="sort"
                                value="word-ascending"
                                defaultChecked
                            />
                                Word — Ascending
                        </label>

                        <label>
                            <input
                                type="radio"
                                name="sort"
                                value="word-descending"
                            />
                                Word — Descending
                        </label>
                    </div>
                </section>


                {/* Dictionary Catalog */}
                <section
                    className="dictionary-catalog"
                    aria-label="Dictionary entries"
                >
                    {isLoading ? (
                        <p>Loading Dictionary...</p>
                    ) : (
                        dictionaryContent.map((data) => {
                            const isExpanded = expandedCard === data.senseId;

                            return(
                                <article 
                                    className={`dictionary-card ${isExpanded ? "expanded" : ""}`}
                                    key={data.senseId}
                                >
                                    <button
                                        type="button"
                                        className="dictionary-card-header"
                                        aria-expanded={isExpanded}
                                        onClick={() => handleCardToggle(data.senseId)}
                                    >
                                        <span className="dictionary-word">
                                            {data.word}
                                        </span>

                                        <span className="dictionary-type">
                                            {data.type}
                                        </span>
                                    </button>

                                    <div className="dictionary-card-content">
                                        <div className='dictionary-card-content-inner'>
                                            <section className="dictionary-detail">
                                                <h2>Translation</h2>
                                                <p>{data.translations.join(", ")}</p>
                                            </section>

                                            <section className="dictionary-detail">
                                                <p>{data.synonyms.join(", ")}</p>
                                                <h2>Synonyms</h2>
                                            </section>

                                            <section className="dictionary-detail">
                                                <h2>Description</h2>

                                                <p className='dictionary-detail-description'>{data.description || "-"}</p>
                                            </section>
                                
                                            <button
                                                type="button"
                                                className="dictionary-card-close"
                                                onClick={handleCardClose}
                                            >
                                                Close
                                            </button>
                                        </div>
                                    </div>
                                </article>
                            );
                        })
                    )}
                </section>
                
                {/* Infinite Scroll Sentinel */}
                <div 
                    ref={observerTarget}
                    className='dictionary-sentinel'
                    aria-hidden="true"
                />

                {isFetchingNextPage && (
                    <p className='dictionary-loading'>
                        Loading more words...
                    </p>
                )}

                {!isLoading &&
                    dictionaryContent.length === 0 && (
                        <p className='dictionary-empty'>
                            No dictionary entries found.
                        </p>
                )}
            </main>
        </UserLayout>
    );
};