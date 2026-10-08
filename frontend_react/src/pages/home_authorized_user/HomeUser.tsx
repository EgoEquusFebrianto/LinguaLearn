import { FaBook, FaGamepad } from "react-icons/fa";
import { Link } from "react-router-dom";

import "./HomeUser.css";

export const HomeUser = () => {

    return (
        <div className="home-user">
            <section className="home-welcome">
                <h1>Selamat datang kembali! 👋</h1>

                <p>Mari lanjutkan perjalanan belajar bahasa Inggris Anda.</p>
            </section>

            <section className="home-quick-actions">
                <Link
                    to="/dictionary"
                    className="quick-action-card"
                >
                    <FaBook />

                    <div>
                        <h2>Kamus</h2>
                        <p>Cari arti dan informasi kata bahasa Inggris.</p>
                    </div>
                </Link>

                <Link
                    to="/games"
                    className="quick-action-card"
                >
                    <FaGamepad />

                    <div>
                        <h2>Challenge</h2>
                        <p>Uji kemampuan bahasa Inggris Anda melalui game.</p>
                    </div>
                </Link>
            </section>

            <section className="learning-section">
                <div className="section-header"><h2>Progress Belajar</h2></div>

                <div className="progress-grid">
                    <div className="progress-card">
                        <span>Games Dimainkan</span>
                        <strong>12</strong>
                    </div>

                    <div className="progress-card">
                        <span>Jawaban Benar</span>
                        <strong>38</strong>
                    </div>

                    <div className="progress-card">
                        <span>Total Poin</span>
                        <strong>420</strong>
                    </div>

                </div>
            </section>

            <section className="learning-section">
                <div className="section-header">
                    <h2>Aktivitas Terakhir</h2>
                </div>

                <div className="activity-list">

                    <div className="activity-item">
                        <span>Matching Words</span>
                        <strong>80%</strong>
                    </div>

                    <div className="activity-item">
                        <span>Guess Translation</span>
                        <strong>70%</strong>
                    </div>

                    <div className="activity-item">
                        <span>Sentence Ordering</span>
                        <strong>90%</strong>
                    </div>
                </div>
            </section>
        </div>
    );
};