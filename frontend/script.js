document.addEventListener('DOMContentLoaded', () => {
    const authSection = document.getElementById('auth-section');
    const movieSection = document.getElementById('movie-section');

    const registerForm = document.getElementById('register-form');
    const loginForm = document.getElementById('login-form');
    const addMovieForm = document.getElementById('add-movie-form');
    const logoutButton = document.getElementById('logout-button');

    const registerMessage = document.getElementById('register-message');
    const loginMessage = document.getElementById('login-message');
    const addMovieMessage = document.getElementById('add-movie-message');
    const movieNameInput = document.getElementById('movie-name');
    const movieUl = document.querySelector('#movie-list ul');

    const NODE_BACKEND_URL = 'http://localhost:3100'; // Node.js backend
    const GO_BACKEND_URL = 'http://localhost:4100'; // Go backend runs on 4100

    function showMessage(element, message, type) {
        element.textContent = message;
        element.className = `message ${type}`;
        element.classList.remove('hidden');
        setTimeout(() => {
            element.classList.add('hidden');
        }, 5000);
    }

    function checkAuth() {
        const token = localStorage.getItem('jwtToken');
        if (token) {
            authSection.classList.add('hidden');
            movieSection.classList.remove('hidden');
            fetchMovies();
        } else {
            authSection.classList.remove('hidden');
            movieSection.classList.add('hidden');
        }
    }

    async function fetchMovies() {
        const token = localStorage.getItem('jwtToken');
        if (!token) {
            return;
        }

        try {
            const response = await fetch(`${GO_BACKEND_URL}/movies`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                if (response.status === 401) {
                    // Token expired or invalid, log out
                    localStorage.removeItem('jwtToken');
                    checkAuth();
                    showMessage(loginMessage, 'Session expired. Please log in again.', 'error');
                }
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const movies = await response.json();
            console.log('Movies received:', movies);
            movieUl.innerHTML = ''; // Clear existing list
            if (movies.length === 0) {
                movieUl.innerHTML = '<li>No favorite movies yet. Add some!</li>';
            } else {
                movies.forEach(movie => {
                    const li = document.createElement('li');
                    li.textContent = movie.movie_name;
                    movieUl.appendChild(li);
                });
            }
        } catch (error) {
            console.error('Error fetching movies:', error);
            showMessage(addMovieMessage, 'Failed to load movies.', 'error');
        }
    }

    registerForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const email = e.target.elements['register-email'].value;
        const password = e.target.elements['register-password'].value;

        try {
            const response = await fetch(`${NODE_BACKEND_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, password })
            });

            const data = await response.json();
            if (response.ok) {
                showMessage(registerMessage, data.message, 'success');
                e.target.reset();
            } else {
                showMessage(registerMessage, data.message || 'Registration failed', 'error');
            }
        } catch (error) {
            console.error('Error during registration:', error);
            showMessage(registerMessage, 'Network error or server unreachable', 'error');
        }
    });

    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const email = e.target.elements['login-email'].value;
        const password = e.target.elements['login-password'].value;

        try {
            const response = await fetch(`${NODE_BACKEND_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, password })
            });

            const data = await response.json();
            if (response.ok) {
                localStorage.setItem('jwtToken', data.token);
                showMessage(loginMessage, data.message, 'success');
                e.target.reset();
                checkAuth(); // Update UI after login
            } else {
                showMessage(loginMessage, data.message || 'Login failed', 'error');
            }
        } catch (error) {
            console.error('Error during login:', error);
            showMessage(loginMessage, 'Network error or server unreachable', 'error');
        }
    });

    addMovieForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const movieName = movieNameInput.value.trim();
        if (!movieName) {
            showMessage(addMovieMessage, 'Movie name cannot be empty.', 'error');
            return;
        }

        const token = localStorage.getItem('jwtToken');
        if (!token) {
            showMessage(addMovieMessage, 'Not logged in.', 'error');
            return;
        }

        try {
            const response = await fetch(`${GO_BACKEND_URL}/movies`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ movie_name: movieName })
            });

            if (!response.ok) {
                if (response.status === 401) {
                    localStorage.removeItem('jwtToken');
                    checkAuth();
                    showMessage(addMovieMessage, 'Session expired. Please log in again.', 'error');
                }
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const updatedMovies = await response.json();
            movieUl.innerHTML = ''; // Clear existing list
            if (updatedMovies.length === 0) {
                movieUl.innerHTML = '<li>No favorite movies yet. Add some!</li>';
            } else {
                updatedMovies.forEach(movie => {
                    const li = document.createElement('li');
                    li.textContent = movie.movie_name;
                    movieUl.appendChild(li);
                });
            }
            showMessage(addMovieMessage, 'Movie added successfully!', 'success');
            movieNameInput.value = ''; // Clear input
        } catch (error) {
            console.error('Error adding movie:', error);
            showMessage(addMovieMessage, 'Failed to add movie.', 'error');
        }
    });

    logoutButton.addEventListener('click', () => {
        localStorage.removeItem('jwtToken');
        showMessage(loginMessage, 'Logged out successfully.', 'success');
        checkAuth(); // Update UI after logout
    });

    // Initial check on page load
    checkAuth();
});
