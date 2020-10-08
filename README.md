### Пропозиція проекту.
Аспект пропозиції | Суть аспекту пропозиції
:---              | :---
Назва програмної системи | Fair Online Judge
Призначення системи. Порівняння з існуючими аналогічними системами. Ефект від її впровадження | <p>Основною задачею розробленої системи (тут і далі наз. **ОЗ**), у противагу до прикладної її складової, яка використовує та будує інфраструктуру навколо **ОЗ**), є визначення близькості текстових файлів (зокрема програмних кодів) у структурному та семантичному розумінні, візуалізація означеної близькості.</p><p>Інфраструктурна складова проекту використоує розроблену технологію **ОЗ** у задачі автоматичного оцінювання близькості (у означеному вище розумінні) текстових даних загальної форми (тут і далі **ІЗ**). Сценарій користування наступний: <ul><li>Надати довільну кількість текстових файлів для подальшої обробки</li><li>Отримати діаграму близькості завантажених корпусів тексту за допомогою розробленого у рамках **ОЗ** рішення,</li><li>Детально вивчити різницю між двома підозрілими на схожість корпусами тексту через веб-інтерфейс</li></ul></p><p>**ІЗ** та має бути вирішена і розгорнена у вигляді веб-сервера.</p></p>Авторам не відомий сервіс, що виконує задачі **ОЗ**, **ІЗ** разом або окремо</p><p>Задача **ОЗ** є дослідницьким напрямком, у якому state-of-the-art технології продовжуюсть з'являтися. Тому інтегрування таких технологій у прикладні рішення має цінність для стимулювання розвитку цієї галузі.</p>
Загальна характеристика задачі, що вирішується |<ul><li>Пошук близькості між текстами може базуватися на статистичних та ін. методах розпізнавання та обробки людської мови (NLP) - насамперед TF-IDF, Word Embedding.</li><li>Невідомо, чи можемо ми модифікувати алгоритми пошуку структурної подібності (т.з. жадібні diff-алгоритми) для якісного вирішення задачі **ОЗ**. Частково мета даної роботи полягає в перевірці такого припущення.</li><li>Візуалізація близькості текстових корпусів може бути виконана у вигляді а) теплових карт (зображення 3D на площині із розподілом насиченості кольору на ній), б) як звичайне двовимірне зображення та в) як повноцінне 3D-зображення.</li><li>Проектування, розробка та розгортання веб-серверу, який виконував би задачу **ІЗ**, є необхідною частиною роботи для демонстрації успіху вирішення задачі **ОЗ**</li></ul>
Мотивація вибору задачі |<ul><li>Дослідницька роботу у напрямку **ОЗ**.</li><li>Використання розв'язків загально-теоретичних задач у прикладній області призведе до більшої зацікавленості до досліджень у напрямку методів розуміння семантики текстових даних.</li><li>Бажання вирішення проблеми справедливого оцінювання робіт шляхом візуалізації подібності рішень.</li></ul>
Аналіз здійсненності. Необхідні для виконання проекту ресурси. Основний ризик для проекту | Складність полягає у плануванні розробки проекту, тобто рішенні підзадач, зазначених у відповідних пуктах, проектуванні архітектури, опануванні фронт-енд частини. Ризиком є можливість неіснування на даний момент задовільного за часом виконання та якістю роботи вирішення задачі **ОЗ**; можлива недостатність часу для виконання бажаного обсягу роботи, що призведе до урізання функціоналу; вибрана ітеративна модель, що може сповільнити процес розробки; відсутність розробника, що на достатньому рівні вміє робити Front-End.
Модель виробничого циклу. Обгрунтування вибору. Відмінні риси процесу розробки обраної системи | В якості моделі була вибрана ітеративна модель. З вибором такої моделі на кожному кроці буде змога повенутися на попередні для редагування аспектів проекту.

Команда: Геворгян Артем (розробка **ОЗ** та **ІЗ**, Пилипець Гліб (розробка **ОЗ** та **ІЗ**), Сизько Микола (розробка фронтенду **ІЗ** та розгортання серверу)
Провідний програміст Геворгян Артем
