import { useEffect, useState } from 'react';
import payload from './data_example.json';
import "./UserDictionary.css";
import { UserLayout } from '../../layouts/user/UserLayout';

type Sense = {
    translations: string[];
    synonyms: string[];
    type: string;
    description: string;
}

type DictionaryEntry = {
    id: string;
    word_uuid: string;
    word: string;
    senses: Sense[];
}

type DictionaryItem = {
    id: string;
    senseId: string;
    word: string;
    type: string;
    translations: string[];
    synonyms: string[];
    description: string;
}

export const UserDictionary = () => {
    const [datas, setData] = useState<DictionaryItem[]>([]);
    const [activeTab, setActiveTab] = useState("all");
    const [isSortOpen, setIsSortOpen] = useState(false);
    const [expandedCard, setExpandedCard] = useState<string | null>(null);

    const normalize = () => {
        const res: DictionaryItem[] = payload.flatMap(
            ({id, word, senses}) => senses.map(
                (sense, idx) => ({
                    senseId: `${id}-${idx}`,
                    id,
                    word,
                    type: sense.type,
                    translations: sense.translations ?? [],
                    synonyms: sense.synonyms ?? [],
                    description: sense.description ?? "",
                })
            )
        )

        setData(res);
    };

    useEffect(() => {
        const fetchData = async () => {
            await normalize();
        }

        fetchData();
    }, [])

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

    console.log("Active tabs=", activeTab)

    return (
        <UserLayout>
            <main className="dictionary-page">
                {/* Page Header */}
                <header className="dictionary-header">
                    <h1>My Dictionary</h1>
                </header>

                {/* Dictionary Tabs */}
                <nav className="dictionary-tabs" aria-label="Dictionary filters">
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
                    {datas.map((data) =>
                        {
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
                                                <h2>Synonyms</h2>
                                                <p>{data.synonyms.join(", ")}</p>
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
                            )
                        }
                    )}
                </section>
            </main>
        </UserLayout>
    );
};