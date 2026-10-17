# Justificación del Lenguaje: Go

Para la implementación de la lógica de dominio (basada en la **HU001**), se ha seleccionado **Go** como lenguaje de programación por los siguientes motivos técnicos y de diseño:

1. **Tipos propios enfocados a la HU001:** Como Jaime necesita buscar calles con sombra, Go nos viene genial para crear tipos propios con campos privados (ej: `Coordenada`). Desde fuera del paquete solo se puede crear con NuevaCoordenada, así que no puede existir una coordenada fuera de rango, haciendo el modelado mucho más seguro.
2. **Inmutabilidad de los Objetos Valor:** En Go es bastante directo lograr la inmutabilidad: podemos poner los atributos en minúscula para que sean privados y no ofrecer ningún método que los modifique. Así nos aseguramos de que una vez creado el dato, nadie lo modifique desde fuera.
3. **Gestión de errores sin excepciones:** Como Go te hace devolver el error junto con el resultado, nos va a obligar a validar siempre si los datos de entrada tienen sentido (por ejemplo, al intentar crear una coordenada) antes de seguir.
4. **Tipado estático y composición:** Al ser de tipado estático, el compilador nos va a decir muchos fallos de tipos antes de tener que probar nada a mano. Además, usa composición en lugar de herencia clásica de clases, lo que nos encaja para ir montando las entidades paso a paso. El código del dominio está en `internal/dominio`, y en Go la carpeta `internal` impide que otros módulos lo importen.
