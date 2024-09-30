Claro, aquí tienes una descripción detallada para crear un diagrama de flujo del código en Go que has proporcionado. Te proporcionaré la secuencia de pasos y cómo deben conectarse, así como el texto que puedes usar en cada parte.
Elementos del Diagrama de Flujo

    Inicio: Representado por un círculo negro.
    Proceso: Representado por un rectángulo.
    Decisión: Representado por un rombo.
    Fin: Representado por un círculo negro con borde blanco.

Pasos para el Diagrama de Flujo
1. Inicio

    Forma: Círculo negro
    Texto: "Inicio"

2. Llamar a sum(1, 2)

    Forma: Rectángulo
    Texto: "Llamar a sum(1, 2)"

3. En la función sum

    Forma: Rectángulo
    Texto: "Imprimir nums"

4. Inicializar total

    Forma: Rectángulo
    Texto: "total <- 0"

5. Bucle: Para cada num en nums

    Forma: Rombo
    Texto: "¿Hay más números en nums?"

6. Sumar num a total

    Forma: Rectángulo
    Texto: "total <- total + num"

7. Imprimir total

    Forma: Rectángulo
    Texto: "Imprimir total"

8. Fin de sum

    Forma: Círculo negro con borde blanco
    Texto: "Fin de sum"

9. Llamar a sum(1, 2, 3)

    Forma: Rectángulo
    Texto: "Llamar a sum(1, 2, 3)"

10. Llamar a sum(nums...)

    Forma: Rectángulo
    Texto: "Crear array nums = [1, 2, 3, 4]"

11. Fin del programa

    Forma: Círculo negro con borde blanco
    Texto: "Fin del programa"

Estructura del Diagrama de Flujo

Aquí tienes la secuencia completa con conexiones entre los elementos:

    Inicio →
    Llamar a sum(1, 2) →
    Imprimir nums →
    Inicializar total →
    ¿Hay más números en nums? →
        Sí →
            Sumar num a total → Regresa al paso 5.
        No →
            Imprimir total →
            Fin de sum → Regresa al paso 9.
    Llamar a sum(1, 2, 3) →
        Repite el proceso desde Imprimir nums.
    Crear array nums = [1, 2, 3, 4] →
    Llamar a sum(nums...) →
        Repite el proceso desde Imprimir nums.
    Fin del programa