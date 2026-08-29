import { useEffect, useRef } from 'react';

export const HeroSectionPublic = () => {
  const heroRef = useRef(null);

  useEffect(() => {
    const handleScroll = () => {
      if (heroRef.current) {
        const scrolled = window.pageYOffset;
        heroRef.current.style.transform = `translateY(${scrolled * 0.5}px)`;
        heroRef.current.style.opacity = 1 - scrolled / 800;
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  return (
    <section className="hero-section" ref={heroRef}>
      <div className="hero-background">
        <div className="stars">
          {[...Array(50)].map((_, i) => (
            <div 
              key={i} 
              className="star" 
              style={{
                left: `${Math.random() * 100}%`,
                top: `${Math.random() * 100}%`,
                animationDelay: `${Math.random() * 3}s`,
                width: `${Math.random() * 3 + 1}px`,
                height: `${Math.random() * 3 + 1}px`,
              }}
            />
          ))}
        </div>
        <div className="shooting-stars">
          {[...Array(3)].map((_, i) => (
            <div 
              key={i} 
              className="shooting-star"
              style={{
                animationDelay: `${i * 4 + 2}s`,
                top: `${20 + i * 15}%`,
              }}
            />
          ))}
        </div>
      </div>

      <div className="hero-content">
        <div className="hero-text">
          <span className="hero-badge">🌟 Belajar Jadi Menyenangkan</span>
          <h1 className="hero-title">
            <span className="title-highlight">Kuasai</span> Bahasa Inggris
            <br />
            <span className="title-gradient">Dengan Cara Baru</span>
          </h1>
          <p className="hero-description">
            Gabungkan kamus interaktif dan 3 permainan seru untuk mengasah 
            kosakata, grammar, dan kemampuan terjemahanmu.
          </p>
          <div className="hero-buttons">
            <button className="hero-btn-primary">
              Mulai Belajar
              <span className="btn-icon">→</span>
            </button>
            <button className="hero-btn-secondary">
              Lihat Fitur
            </button>
          </div>
          <div className="hero-trust">
            <div className="trust-avatars">
              <img src="https://i.pravatar.cc/40?img=1" alt="user" />
              <img src="https://i.pravatar.cc/40?img=2" alt="user" />
              <img src="https://i.pravatar.cc/40?img=3" alt="user" />
              <img src="https://i.pravatar.cc/40?img=4" alt="user" />
              <div className="trust-more">+5K</div>
            </div>
            <span className="trust-text">Bergabung dengan 10.000+ pelajar</span>
          </div>
        </div>

        <div className="hero-illustration">
          <div className="floating-card card-1">
            <span>📖</span>
            <span>Kamus</span>
          </div>
          <div className="floating-card card-2">
            <span>🎮</span>
            <span>Game</span>
          </div>
          <div className="floating-card card-3">
            <span>🏆</span>
            <span>Pencapaian</span>
          </div>
          <div className="hero-circle"></div>
          <div className="hero-circle-2"></div>
        </div>
      </div>
    </section>
  );
}