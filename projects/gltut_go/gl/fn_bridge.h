
#ifndef _fn_bridge_h_
#define _fn_bridge_h_

#include <GL/gl.h>

// GLAPI void APIENTRY glGenBuffers (GLsizei n, GLuint *buffers);
static inline void fn_bridge_void_GLsizei_GLuint_p(void (*fn)(GLsizei, GLuint *), GLsizei a1, GLuint *a2)
{
    fn(a1, a2);
}

#endif // _fn_bridge_h_
