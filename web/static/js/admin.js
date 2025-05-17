document.addEventListener('DOMContentLoaded', initAdmin);

async function initAdmin() {
    await loadContentList();
    setupEventListeners();
}

async function loadContentList() {
    try {
        const response = await fetch('/api/content');
        if (!response.ok) throw new Error('Failed to load content');
        
        const { data } = await response.json();
        const container = document.getElementById('content-list');
        
        container.innerHTML = data.map(content => `
            <div class="bg-gray-50 p-4 rounded-lg mb-2">
                <h3 class="font-bold">${content.title}</h3>
                <p class="text-sm text-gray-600">${content.body.substring(0, 80)}...</p>
                <div class="mt-2 space-x-2">
                    <button onclick="editContent('${content.id}')" class="text-blue-600 hover:text-blue-800">Edit</button>
                    <button onclick="deleteContent('${content.id}')" class="text-red-600 hover:text-red-800">Delete</button>
                </div>
            </div>
        `).join('');
    } catch (error) {
        showError('Failed to load content: ' + error.message);
    }
}

function setupEventListeners() {
    document.getElementById('new-content-btn').addEventListener('click', () => {
        window.location.href = '/admin/content/new';
    });
}

async function deleteContent(id) {
    if (!confirm('Are you sure you want to delete this content?')) return;

    try {
        const response = await fetch(`/api/content/${id}`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        if (!response.ok) throw new Error('Delete failed');
        await loadContentList();
    } catch (error) {
        showError('Delete failed: ' + error.message);
    }
}

function showError(message) {
    const errorDiv = document.createElement('div');
    errorDiv.className = 'bg-red-100 border-red-400 text-red-700 px-4 py-3 rounded mb-4';
    errorDiv.innerHTML = `
        <strong>Error:</strong>
        <span class="block sm:inline">${message}</span>
    `;
    document.getElementById('content-list').prepend(errorDiv);
}