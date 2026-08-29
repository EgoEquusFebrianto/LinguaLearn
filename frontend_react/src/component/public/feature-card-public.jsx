// src/components/FeatureCard.jsx
import React, { useState } from 'react';

export const FeatureCardPublic = ({ feature }) => {
  const [isHovered, setIsHovered] = useState(false);

  return (
    <div 
      className={`feature-card ${isHovered ? 'hovered' : ''}`}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      style={{
        '--card-gradient': feature.gradient,
      }}
    >
      <div className="feature-card-glow" style={{ background: feature.gradient }}></div>
      
      <div className="feature-card-icon" style={{ color: feature.color }}>
        {feature.icon}
      </div>
      
      <h3 className="feature-card-title">{feature.title}</h3>
      <p className="feature-card-description">{feature.description}</p>
      
      <button className="feature-card-btn">
        <span>Pelajari</span>
        <span className="btn-arrow">→</span>
      </button>
      
      <div className="feature-card-number">#{String(feature.id).padStart(2, '0')}</div>
    </div>
  );
};