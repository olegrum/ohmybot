# About
This is our pet project that gonna be multiplatform bot constructor. Nowbody knows if this project ever been finished.
# The task :D
The task was formulated by our friend as follows:
```
Олень, [22.07.21 02:28]
Итак, ТЗ "в общих чертах"

Общая идея: нужна приложуха, с помощью которой можно написать алгоритмы для чат-бота с поддержкой кнопок и присобачивать туда любые апишки, то есть реализовывать в любых системах, где возможно делать чат-бота с кнопками. Самые приоритетные - ВК, Телеграм, Дискорд. В принципе этого достаточно, позже можно будет при желании чисто по фану прицепить вотсапп, стим и пр., можно будет провести опрос и понять насколько это вообще необходимо.
Информация, собираемая с чат-бота, должна где-нибудь храниться в одном месте, и при обращении из чат-бота на любом клиенте её можно будет получить в обозначенном алгоритмами формате. Также должна быть реализована возможность автоматических рассылок, которые администратор готовит заранее.

Теперь чуть подробнее.

Админка
Должна быть запилена возможность администрирования всего этого, то есть изменение внутренних алгоритмов чат-бота. После их редактирования чат-бот на всех подключенных приложениях будет обновлён. Админку можно сделать с веб-мордой, а можно попробовать через сами чат-боты и реализовать, но это наверное будет довольно неудобно, такой консольный режим получится, лучше с возможностью визуализации. Рекомендую посмотреть как щас сделано в смартботе, могу показать очно или по интернету, как всё выглядит у меня.

Кросс-платформенность
Скорее всего потребуется механизм какой-нибудь внутренней регистрации, чтобы чат-бот мог определить, что вот эти вот аккаунты в ВК и в телеге - один и тот же человек. Чтобы он мог, например, записаться на настолки через ВК, потом написать в бот в телеге, а тот уже знает, что этот человек записан на игры. Правда, придётся как-то заставить всех там регистрироваться, но это уже мои проблемы) На крайняк буду вручную обозначать в админке, что вот эти вот аккаунты - один человек. Вместе с этим можно также будет запилить подписку на рассылки в определённых приложениях. Например, у меня вот есть ВК и телега, но хочу чтобы бот мне писал рассылки только в телегу, но если я что-то напишу с ВК - он меня узнает и откликнется соответственно.
Кроме того, наверняка функционал на разных платформах будет разный. Где-то можно прилепить опрос к сообщению, где-то нельзя, где-то можно песенки или картинки, где-то может и нет. По возможности нужно запилить так, чтобы поддерживалось максимум возможностей от каждой платформы, но чтобы если я например делаю опрос, а он хуяк и не поддерживается в дискорде - мне предупреждение показывалось, что опросы не будут работать в дискорде. Возможно даже запилить возможность части алгоритма делать по-разному в зависимости от приложения.

Алгоритмы
Там нужны всякие обработчики событий. Прислали сообщение или нажали на кнопку - какая-нибудь реакция. Должны быть переходы к тем или иным частям алгоритма, таймеры, возможно, парсеры. Должна быть реализована булева логика для проверки всяких условий, чтобы выдавать разные реакции в зависимости от того, что написал пользак, какой это пользак, какие у него значения в базе записаны, какие значения других переменных и т.д. Опять же, могу показать, как это всё сейчас запилено в смартботе и какой там есть функционал. Само собой должны быть обращения к базе, где всякая инфа будет записываться и обновляться. В смартботе также недостаток функционала компенсируется своим языком, на котором можно какие-нибудь скрипты накатать.

Олень, [22.07.21 02:28]
База
Будут глобальные переменные и локальные. В данном случае "переменная" - это какая-то запись в БД, отображающая состояние чего-то или какую-нибудь собранную от пользака или от админа инфу. Очень хорошо если будет простой "консольный" (через самого чат-бота) вариант поменять значение любой переменной. В переменных ожидаются текст/числа/булевы, а также различные массивы, массивы массивов и мб другие структуры из этих же типов. Можно запилить и всякую поддержку файлов, картинок, по крайней мере хорошо бы иметь возможность например делать рассылку или просто присылать сообщения с выбранными прилепленными изображениями.
Насчёт базы настолок пока хз, думаю я таки доделаю её в ноушне, потому что просматривать её всё равно будет удобнее там или в какой-нибудь ещё веб-морде, нежели в чате, и нужно будет её прилепить по его АПИшке к боту. Вопрос требует изучения короче.
Где и как всё это хранить - отдельный вопрос, у меня так-то есть сервак синоложи на юниксе, для личного использования мощностей наверное хватит. Если же продвинуть эту штуку для дальнейшего использования в более широкой аудитории, то вероятно стоит реализовать возможность указать свой сервак, сделать это в виде сервиса там, например.

Рассылки
Автоматизация, которая в указанное время или по иному триггеру пришлёт пользаку указанную админом информацию. При этом "прогресс" общения с ботом должен сохраняться, то есть входящее сообщение от бота не должно прерывать цепочку взаимодействия с ботом. Например: сижу я в боте, читаю инфу, делаю там поиск по настолкам, например, и тут мне приходит рассылка с приглашением на игротеку. Должна быть возможность вернуться оттуда назад, в то же меню, где я общался с ботом.

Посты
Если получится ещё и прилепить автоматическое создание постов в группе ВК наравне с рассылкой от бота - будет вообще збс.
```