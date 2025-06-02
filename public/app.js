// MindShelf JavaScript
// A minimalist link management system built with Alpine.js

function app() {
    return {
        // Auth state
        isAuthenticated: false,
        authTab: 'login',
        token: null,
        loginForm: {
            email: '',
            password: ''
        },
        registerForm: {
            email: '',
            password: ''
        },

        // Links
        links: [],
        isLoading: false,
        searchTerm: '',

        // Link form
        showAddLinkModal: false,
        editingLink: null,
        linkForm: {
            url: '',
            title: '',
            description: '',
            tags: []
        },
        tagsInput: '',

        // Delete modal
        showDeleteModal: false,
        linkIdToDelete: null,

        // Notifications
        notyf: null,

        // Initialize
        init() {
            // Initialize notifications
            this.notyf = new Notyf({
                duration: 3000,
                position: { x: 'right', y: 'top' },
                types: [
                    {
                        type: 'success',
                        background: '#3b82f6',
                    },
                    {
                        type: 'error',
                        background: '#ef4444',
                    }
                ]
            });
            
            // Add event listeners for keyboard shortcuts
            window.addEventListener('keydown', (e) => this.handleKeyDown(e));
            
            // Add URL input event listener for auto metadata extraction
            // We'll use Alpine's x-effect for this in the HTML
        },

        // Auth methods
        checkAuth() {
            this.token = localStorage.getItem('token');
            this.isAuthenticated = !!this.token;
            
            if (this.isAuthenticated) {
                this.fetchLinks();
            }
        },

        async login() {
            this.isLoading = true;
            try {
                const response = await fetch('/api/auth/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(this.loginForm)
                });

                if (!response.ok) {
                    const error = await response.text();
                    throw new Error(error || 'Invalid credentials');
                }

                const data = await response.json();
                this.token = data.token;
                localStorage.setItem('token', this.token);
                this.isAuthenticated = true;
                this.notyf.success('Login successful');
                this.fetchLinks();
                this.resetLoginForm();
            } catch (error) {
                this.notyf.error(error.message || 'Login failed');
            } finally {
                this.isLoading = false;
            }
        },

        async register() {
            this.isLoading = true;
            try {
                const response = await fetch('/api/auth/register', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(this.registerForm)
                });

                if (!response.ok) {
                    const error = await response.text();
                    throw new Error(error || 'Registration failed');
                }

                this.notyf.success('Registration successful! You can now login.');
                this.authTab = 'login';
                this.resetRegisterForm();
            } catch (error) {
                this.notyf.error(error.message || 'Registration failed');
            } finally {
                this.isLoading = false;
            }
        },

        logout() {
            localStorage.removeItem('token');
            this.isAuthenticated = false;
            this.token = null;
            this.links = [];
            this.notyf.success('Logged out successfully');
        },

        resetLoginForm() {
            this.loginForm = { email: '', password: '' };
        },

        resetRegisterForm() {
            this.registerForm = { email: '', password: '' };
        },

        // Links methods
        async fetchLinks() {
            this.isLoading = true;
            try {
                const response = await fetch('/api/links', {
                    headers: { 
                        'Authorization': `Bearer ${this.token}` 
                    }
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        this.logout();
                        throw new Error('Session expired. Please login again.');
                    }
                    throw new Error('Failed to fetch links');
                }

                this.links = await response.json();
            } catch (error) {
                this.notyf.error(error.message);
            } finally {
                this.isLoading = false;
            }
        },

        async searchLinks() {
            if (!this.searchTerm.trim()) {
                this.fetchLinks();
                return;
            }

            this.isLoading = true;
            try {
                const response = await fetch(`/api/links/search?q=${encodeURIComponent(this.searchTerm)}`, {
                    headers: { 
                        'Authorization': `Bearer ${this.token}` 
                    }
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        this.logout();
                        throw new Error('Session expired. Please login again.');
                    }
                    throw new Error('Search failed');
                }

                this.links = await response.json();
            } catch (error) {
                this.notyf.error(error.message);
            } finally {
                this.isLoading = false;
            }
        },

        searchByTag(tag) {
            this.searchTerm = tag;
            this.searchLinks();
        },

        openAddLinkModal() {
            this.resetLinkForm();
            this.editingLink = null;
            this.showAddLinkModal = true;
        },

        editLink(link) {
            this.editingLink = link;
            this.linkForm = {
                url: link.url,
                title: link.title || '',
                description: link.description || '',
                tags: link.tags || []
            };
            this.tagsInput = (link.tags || []).join(', ');
            this.showAddLinkModal = true;
        },

        resetLinkForm() {
            this.linkForm = {
                url: '',
                title: '',
                description: '',
                tags: []
            };
            this.tagsInput = '';
        },

        async addLink() {
            // Process tags
            this.processTagsInput();
            
            this.isLoading = true;
            try {
                const response = await fetch('/api/links', {
                    method: 'POST',
                    headers: { 
                        'Authorization': `Bearer ${this.token}`,
                        'Content-Type': 'application/json' 
                    },
                    body: JSON.stringify(this.linkForm)
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        this.logout();
                        throw new Error('Session expired. Please login again.');
                    }
                    throw new Error('Failed to add link');
                }

                this.showAddLinkModal = false;
                this.resetLinkForm();
                this.notyf.success('Link added successfully');
                this.fetchLinks();
            } catch (error) {
                this.notyf.error(error.message);
            } finally {
                this.isLoading = false;
            }
        },

        async updateLink() {
            // Process tags
            this.processTagsInput();
            
            this.isLoading = true;
            try {
                const response = await fetch(`/api/links/${this.editingLink.id}`, {
                    method: 'PUT',
                    headers: { 
                        'Authorization': `Bearer ${this.token}`,
                        'Content-Type': 'application/json' 
                    },
                    body: JSON.stringify(this.linkForm)
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        this.logout();
                        throw new Error('Session expired. Please login again.');
                    }
                    throw new Error('Failed to update link');
                }

                this.showAddLinkModal = false;
                this.editingLink = null;
                this.resetLinkForm();
                this.notyf.success('Link updated successfully');
                this.fetchLinks();
            } catch (error) {
                this.notyf.error(error.message);
            } finally {
                this.isLoading = false;
            }
        },

        processTagsInput() {
            if (this.tagsInput.trim()) {
                this.linkForm.tags = this.tagsInput.split(',')
                    .map(tag => tag.trim())
                    .filter(tag => tag !== '');
            } else {
                this.linkForm.tags = [];
            }
        },

        confirmDelete(linkId) {
            this.linkIdToDelete = linkId;
            this.showDeleteModal = true;
        },

        async deleteLink() {
            this.isLoading = true;
            const idToDelete = this.linkIdToDelete; // Store id locally before resetting
            
            try {
                const response = await fetch(`/api/links/${idToDelete}`, {
                    method: 'DELETE',
                    headers: { 
                        'Authorization': `Bearer ${this.token}`
                    }
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        this.logout();
                        throw new Error('Session expired. Please login again.');
                    }
                    throw new Error('Failed to delete link');
                }

                this.showDeleteModal = false;
                this.linkIdToDelete = null;
                this.notyf.success('Link deleted successfully');
                this.links = this.links.filter(link => link.id !== idToDelete);
            } catch (error) {
                this.notyf.error(error.message);
            } finally {
                this.isLoading = false;
            }
        },

        // Helper functions
        formatUrl(url) {
            try {
                const urlObj = new URL(url);
                return urlObj.hostname + (urlObj.pathname !== '/' ? urlObj.pathname : '');
            } catch (e) {
                return url;
            }
        },

        formatDate(dateString) {
            if (!dateString) return '';
            const date = new Date(dateString);
            return date.toLocaleDateString('en-US', {
                year: 'numeric',
                month: 'short',
                day: 'numeric'
            });
        },
        
        // Automatically extract title from URL when adding a new link
        async autoFillMetadata() {
            if (!this.linkForm.url || this.linkForm.url.trim() === '') {
                return;
            }
            
            // Only try to extract metadata if title is empty
            if (!this.linkForm.title || this.linkForm.title.trim() === '') {
                this.notyf.success('Fetching website information...');
                try {
                    // Due to CORS restrictions, we would normally need a proxy server.
                    // For simplicity, we'll just use a basic title extraction method
                    const urlObject = new URL(this.linkForm.url);
                    // Generate a title from the domain if we can't fetch it
                    const domain = urlObject.hostname.replace('www.', '');
                    const domainParts = domain.split('.');
                    if (domainParts.length > 0) {
                        this.linkForm.title = domainParts[0].charAt(0).toUpperCase() + domainParts[0].slice(1);
                        
                        // Generate tags from the URL parts
                        const pathParts = urlObject.pathname.split('/').filter(p => p !== '');
                        if (pathParts.length > 0) {
                            const tags = pathParts.slice(0, 2).filter(p => p.length < 20);
                            if (tags.length > 0) {
                                this.tagsInput = tags.join(', ');
                            }
                        }
                    }
                } catch (error) {
                    console.error('Error extracting metadata:', error);
                }
            }
        },
        
        // Handle keyboard shortcuts
        handleKeyDown(event) {
            // Escape key closes modals
            if (event.key === 'Escape') {
                if (this.showAddLinkModal) {
                    this.showAddLinkModal = false;
                }
                if (this.showDeleteModal) {
                    this.showDeleteModal = false;
                }
            }
            
            // Ctrl+/ focuses search box
            if (event.key === '/' && (event.ctrlKey || event.metaKey)) {
                event.preventDefault();
                document.querySelector('.search-bar input')?.focus();
            }
        }
    };
}