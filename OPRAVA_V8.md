# Oprava V8

Opraven opakovaný překlad textu tlačítka Export, který vytvářel řetězec „Exportovatovat…“.

Příčina: překladová vrstva nahrazovala části již přeložených slov a MutationObserver znovu zpracovával vlastní změny.

Oprava:
- zpětný překlad používá pouze přesnou shodu celého textu,
- každý textový uzel si uchovává neměnný původní český zdroj,
- opakované zpracování je idempotentní a text se již nemůže prodlužovat,
- vytvořen nový produkční build.
