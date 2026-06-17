# Optimalizace výkonu – KLIK V7

Hlavní příčinou zpomalování byla původní překladová vrstva. Při každé změně textu nebo otevření dialogu znovu procházela všechny elementy celé stránky. To způsobovalo opakované vykreslování a vysoké zatížení procesoru.

Provedené změny:

1. Překlad se při přepnutí jazyka provede jednou pro aktuální stránku.
2. Pozdější změny se zpracovávají pouze v nově vloženém nebo změněném prvku.
3. Více změn se spojí do jedné dávky přes requestIdleCallback a requestAnimationFrame.
4. Překladové položky jsou setříděné pouze jednou při spuštění aplikace.
5. Zápisy do localStorage jsou sloučené a odložené o 180 ms.
6. Recharts animace byly vypnuté, aby grafy nezatěžovaly hlavní vlákno.
7. Produkční build byl znovu vytvořen a zkontrolován.
