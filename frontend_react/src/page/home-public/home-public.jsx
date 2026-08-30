import { FaBook, FaPuzzlePiece, FaSortAlphaDown, FaKeyboard } from 'react-icons/fa';
import { NavbarPublic } from "../../component/public/navbar-public";
import { FeatureCardPublic } from "../../component/public/feature-card-public";
import { HeroSectionPublic } from "../../component/public/hero-section-public";
import './home-public.css';

export const HomePublic = () => {
  const features = [
    {
      id: 1,
      icon: <FaBook className="feature-icon" />,
      title: 'Kamus Inggris-Indonesia',
      description: 'Lebih dari 50.000 kosakata dengan pengucapan audio dan contoh kalimat interaktif.',
      color: '#FF6B6B',
      gradient: 'linear-gradient(135deg, #FF6B6B 0%, #FF8E53 100%)',
    },
    {
      id: 2,
      icon: <FaPuzzlePiece className="feature-icon" />,
      title: 'Match the Word',
      description: 'Cocokkan kata bahasa Inggris dengan terjemahan bahasa Indonesia-nya. Asah memori visual Anda!',
      color: '#4ECDC4',
      gradient: 'linear-gradient(135deg, #4ECDC4 0%, #44B39D 100%)',
    },
    {
      id: 3,
      icon: <FaSortAlphaDown className="feature-icon" />,
      title: 'Arrange the Sentence',
      description: 'Susun kata-kata acak menjadi kalimat bahasa Inggris yang benar. Latih struktur grammar Anda!',
      color: '#A8E6CF',
      gradient: 'linear-gradient(135deg, #A8E6CF 0%, #7BC8A4 100%)',
    },
    {
      id: 4,
      icon: <FaKeyboard className="feature-icon" />,
      title: 'Translate & Guess',
      description: 'Tebak terjemahan bahasa Inggris dari kata Indonesia melalui pilihan ganda atau mengetik langsung.',
      color: '#FFD93D',
      gradient: 'linear-gradient(135deg, #FFD93D 0%, #F6C90E 100%)',
    },
  ];

  return (
    <div className="home-container">
      <NavbarPublic />
      
      <main className="home-main">
        <HeroSectionPublic />
        
        <section className="features-section">
          <div className="features-header">
            <h2 className="features-title">
              <span className="title-gradient">Fitur Unggulan</span>
            </h2>
            <p className="features-subtitle">
              Kuasai bahasa Inggris dengan 4 cara menyenangkan
            </p>
            <div className="features-divider">
              <span className="divider-line"></span>
              <span className="divider-star">✦</span>
              <span className="divider-line"></span>
            </div>
          </div>

          <div className="features-grid">
            {features.map((feature) => (
              <FeatureCardPublic key={feature.id} feature={feature} />
            ))}
          </div>
        </section>

        <section className="cta-section">
          <div className="cta-content">
            <h3>Mulai Perjalanan Belajarmu</h3>
            <p>Bergabunglah dengan ribuan pelajar lainnya dan kuasai bahasa Inggris dengan cara yang menyenangkan</p>
            <button className="cta-button">
              Daftar Sekarang Gratis
              <span className="cta-arrow">→</span>
            </button>
          </div>
          <div className="cta-stats">
            <div className="stat-item">
              <span className="stat-number">50K+</span>
              <span className="stat-label">Kosakata</span>
            </div>
            <div className="stat-item">
              <span className="stat-number">10K+</span>
              <span className="stat-label">Pengguna Aktif</span>
            </div>
            <div className="stat-item">
              <span className="stat-number">4.9</span>
              <span className="stat-label">Rating Pengguna</span>
            </div>
          </div>
        </section>
      </main>

      <footer className="home-footer">
        <p>© 2026 LinguaLearn. FEBRIANTO KUDADIRI</p>
      </footer>
    </div>
  );
}
