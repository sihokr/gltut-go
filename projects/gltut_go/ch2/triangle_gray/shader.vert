#version 150 core

in vec2 position;
in float color;

out vec3 Color;

void main()
{
    gl_Position = vec4(position, 0.0, 1.0);
    Color = vec3(color, color, color);
}
