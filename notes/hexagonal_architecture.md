## Arquitetura Hexagonal

A arquitetura hexagonal, também conhecida como arquitetura de portos e adaptadores, é um estilo de design de software que visa criar sistemas desacoplados, tornando-os mais flexíveis, testáveis e facilmente adaptáveis a mudanças. Ela foi introduzida por Alistair Cockburn e tem como principal objetivo isolar o núcleo do sistema (a lógica de negócios) das interações externas, como bancos de dados, interfaces de usuário, APIs, etc.

![Draw IO](/github/Draw.png)

### Componentes principais da arquitetura hexagonal:

* **Domínio (Core):** Contém toda a lógica de negócios da aplicação. É a parte mais importante e deve ser independente de detalhes externos, como banco de dados ou frameworks específicos.

* **Portos (Ports):** Representam as interfaces que o núcleo da aplicação expõe ou utiliza. São portas de comunicação entre o domínio e o mundo externo. Os portos podem ser de dois tipos:

    - Portos de entrada (Inbound Ports): Interfaces que definem como o mundo externo interage com o núcleo da aplicação (por exemplo, serviços que consomem a lógica de negócios).

    - Portos de saída (Outbound Ports): Interfaces que definem como o núcleo da aplicação interage com serviços externos (por exemplo, banco de dados, APIs externas).

* **Adaptadores (Adapters):** São implementações que conectam as interfaces das portas com as dependências externas (como UI, banco de dados, APIs, etc.). Há dois tipos de adaptadores:

    - Adaptadores de entrada (Inbound Adapters): Implementam a lógica para conectar interfaces de usuário, APIs externas ou outros clientes à aplicação.
    - Adaptadores de saída (Outbound Adapters): Implementam a comunicação com sistemas externos, como bancos de dados, serviços de mensageria ou APIs.
