import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [tutorials, setTutorials] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchTutorials();
  }, []);

  const fetchTutorials = async () => {
    try {
      const response = await fetch('/api/tutorials');
      const data = await response.json();
      
      if (data.success) {
        setTutorials(data.data);
      } else {
        setError('Failed to load tutorials');
      }
    } catch (err) {
      setError('Error connecting to server');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="App">
      <header className="hero">
        <div className="hero-content">
          <h1 className="title">Welcome to My Tutorials</h1>
          <p className="subtitle">Learn, Build, and Grow Your Skills</p>
        </div>
      </header>

      <main className="main-content">
        <div className="container">
          <h2 className="section-title">Featured Tutorials</h2>
          
          {loading && (
            <div className="loading">Loading tutorials...</div>
          )}
          
          {error && (
            <div className="error">{error}</div>
          )}
          
          {!loading && !error && (
            <div className="tutorials-grid">
              {tutorials.map((tutorial) => (
                <div key={tutorial.id} className="tutorial-card">
                  <div className="card-content">
                    <h3 className="card-title">{tutorial.title}</h3>
                    <p className="card-description">{tutorial.description}</p>
                    <a 
                      href={tutorial.url} 
                      target="_blank" 
                      rel="noopener noreferrer"
                      className="card-link"
                    >
                      Learn More →
                    </a>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </main>

      <footer className="footer">
        <p>&copy; 2026 Tutorial Landing Page. Built with React & Go.</p>
      </footer>
    </div>
  );
}

export default App;
