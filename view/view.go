package view

import (
    "io"
    "fmt"
)

/**************
 * INDEX PAGE *
 **************/

// Builds index page
func ShowIndex(writer io.Writer, posts []map[string]string) {
    args := make(map[string]string)
    args["body"] = PostsToString(posts)
    args["offset"] = ""
    LoadFileWithArgs(writer, "assets/html/index.gohtml", args)
}

/**************
 * BLOG PAGES *
 **************/

// Build a blog page with many posts
func ShowBlog(writer io.Writer, posts []map[string]string, offset int) {
    args := make(map[string]string)
    args["body"] = PostsToString(posts)

    pagination := `<p>`
    if offset >= 1 {
        off := offset-10
        if off < 0 {
            off = 0
        }
        pagination = fmt.Sprintf(`%s<a href="/blog?offset=%d">
            <i class="fa fa-chevron-left" aria-hidden="true"></i>
        </a>`, pagination, off)
    }
    if len(posts) > 10 {
        pagination = fmt.Sprintf(`%s<a href="/blog?offset=%d">
            <i class="fa fa-chevron-right" aria-hidden="true"></i>
        </a>`, pagination, offset+10)
    }
    pagination = fmt.Sprintf("%s</p>", pagination)
    args["offset"] = pagination

    LoadFileWithArgs(writer, "assets/html/blog.gohtml", args)
}

// Displays a blog page with a single page
func ShowPost(writer io.Writer, post map[string]string) {
    args := make(map[string]string)
    args["body"] = PostToString(post, false)
    LoadFileWithArgs(writer, "assets/html/post.gohtml", args)
}

// Provides posts in a machine readable manner
func ProvidePosts(writer io.Writer, posts []map[string]string) {
	// TODO Turn posts into a JSON string
  	// TODO Write JSON payload to writer
}

/****************
 * SUPPORT PAGE *
 ****************/

// Loads support page
func ShowSupport(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/suporte.gohtml")
}

/****************
 * CONTACT PAGE *
 ****************/

// Loads "contact us" page
func ShowContactUs(writer io.Writer) {
    LoadFileWithoutArgs(writer, "assets/html/contato.gohtml")
}
