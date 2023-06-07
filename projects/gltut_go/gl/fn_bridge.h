
#ifndef _fn_bridge_h_
#define _fn_bridge_h_

// #define GLEW_STATIC
// #include <GL/glew.h>
#include <GL/gl.h>
#include <GL/glext.h>

// GLAPI void APIENTRY glGenBuffers (GLsizei n, GLuint *buffers);
static inline void fn_bridge_void_GLsizei_GLuint_p(void (*fn)(GLsizei, GLuint *), GLsizei a1, GLuint *a2)
{
    fn(a1, a2);
}

// GLAPI void APIENTRY glBindBuffer (GLenum target, GLuint buffer);
static inline void fn_bridge_void_GLenum_GLuint(void (*fn)(GLenum, GLuint), GLenum a1, GLuint a2)
{
    fn(a1, a2);
}

// GLAPI void APIENTRY glBufferData (GLenum target, GLsizeiptr size, const void *data, GLenum usage);
static inline void fn_bridge_void_GLenum_GLsizeiptr_void_p_GLenum(void (*fn)(GLenum, GLsizeiptr, const void *, GLenum), GLenum a1, GLsizeiptr a2, const void *a3, GLenum a4)
{
    fn(a1, a2, a3, a4);
}

// GLAPI GLuint APIENTRY glCreateShader (GLenum type);
static inline GLuint fn_bridge_GLuint_GLenum(GLuint (*fn)(GLenum), GLenum a1)
{
    return fn(a1);
}

// GLAPI void APIENTRY glShaderSource (GLuint shader, GLsizei count, const GLchar *const*string, const GLint *length);
static inline void fn_bridge_void_GLuint_GLsizei_GLchar_p_p_GLint_p(
    void (*fn)(GLuint, GLsizei, GLchar **, GLint *),
    GLuint a1, GLsizei a2, GLchar **a3, GLint *a4)
{
    fn(a1, a2, a3, a4);
}

// GLAPI void APIENTRY glCompileShader (GLuint shader);
static inline void fn_bridge_void_GLuint(void (*fn)(GLuint), GLuint a1)
{
    fn(a1);
}

// GLAPI void APIENTRY glGetShaderiv (GLuint shader, GLenum pname, GLint *params);
static inline void fn_bridge_void_GLuint_GLenum_GLint_p(void (*fn)(GLuint, GLenum, GLint *), GLuint a1, GLenum a2, GLint *a3)
{
    fn(a1, a2, a3);
}

// GLAPI void APIENTRY glGetShaderInfoLog (GLuint shader, GLsizei bufSize, GLsizei *length, GLchar *infoLog);
static inline void fn_bridge_void_GLuint_GLsizei_GLsizei_p_GLchar_p(
    void (*fn)(GLuint, GLsizei, GLsizei *, GLchar *), GLuint a1, GLsizei a2, GLsizei *a3, GLchar *a4)
{
    fn(a1, a2, a3, a4);
}

// GLAPI GLuint APIENTRY glCreateProgram (void);
static inline GLuint fn_bridge_GLuint(GLuint (*fn)())
{
    return fn();
}

// GLAPI void APIENTRY glAttachShader (GLuint program, GLuint shader);
static inline void fn_bridge_void_GLuint_GLuint(void (*fn)(GLuint, GLuint), GLuint a1, GLuint a2)
{
    fn(a1, a2);
}

// GLAPI void APIENTRY glBindFragDataLocation (GLuint program, GLuint color, const GLchar *name);
static inline void fn_bridge_void_GLuint_GLuint_GLchar_p(
    void (*fn)(GLuint, GLuint, GLchar *),
    GLuint a1, GLuint a2, GLchar *a3)
{
    fn(a1, a2, a3);
}

// GLAPI GLint APIENTRY glGetAttribLocation (GLuint program, const GLchar *name);
static inline GLint fn_GLint_GLuint_GLchar_p(GLint (*fn)(GLuint, GLchar *), GLuint a1, GLchar *a2)
{
    return fn(a1, a2);
}

// GLAPI void APIENTRY glVertexAttribPointer (GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const void *pointer);
static inline void fn_bridge_void_GLuint_GLint_GLenum_GLboolean_GLsizei_void_p(
    void (*fn)(GLuint, GLint, GLenum, GLboolean, GLsizei, void *),
    GLuint a1, GLint a2, GLenum a3, GLboolean a4, GLsizei a5, void *a6)
{
    fn(a1, a2, a3, a4, a5, a6);
}

// GLAPI GLint APIENTRY glGetUniformLocation (GLuint program, const GLchar *name);
static inline GLint fn_bridge_GLint_GLuint_GLchar_p(
    GLint (*fn)(GLuint, GLchar *),
    GLuint a1, GLchar *a2)
{
    return fn(a1, a2);
}

// GLAPI void APIENTRY glUniform3f (GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
static inline void fn_bridge_void_GLint_GLfloat_GLfloat_GLfloat(
    void (*fn)(GLint, GLfloat, GLfloat, GLfloat),
    GLint a1, GLfloat a2, GLfloat a3, GLfloat a4)
{
    fn(a1, a2, a3, a4);
}

#endif // _fn_bridge_h_
