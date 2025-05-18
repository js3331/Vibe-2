async function submitPost() {
    const content = document.getElementById('postContent').value.trim();
    if (!content) return;

    try {
        const response = await fetch('/api/posts', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ content }),
        });

        if (!response.ok) throw new Error('Failed to create post');

        document.getElementById('postContent').value = '';
        await searchPosts(); // Refresh the posts list
    } catch (error) {
        console.error('Error:', error);
        alert('Failed to create post');
    }
}

async function searchPosts() {
    const query = document.getElementById('searchInput').value;
    try {
        const response = await fetch(`/api/posts?q=${encodeURIComponent(query)}`);
        if (!response.ok) throw new Error('Failed to fetch posts');

        const posts = await response.json();
        displayPosts(posts);
    } catch (error) {
        console.error('Error:', error);
        alert('Failed to fetch posts');
    }
}

function displayPosts(posts) {
    const postsContainer = document.getElementById('posts');
    postsContainer.innerHTML = posts
        .map(post => `
            <div class="post">
                <p>${escapeHtml(post.content)}</p>
            </div>
        `)
        .join('');
}

function escapeHtml(unsafe) {
    return unsafe
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#039;");
}

// Load posts on page load
document.addEventListener('DOMContentLoaded', searchPosts);
