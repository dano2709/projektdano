# Dokumentace řešení KLIK

## Technologie
React, TypeScript, Vite, Recharts a Lucide Icons. Data jsou pro demonstrační účely uložena v paměti prohlížeče. Architektura UI je připravena k rozdělení do modulů a napojení na API.

## Design tokeny
- Primární zelená: `#00843D`
- Tmavá zelená: `#00573F`
- Světle zelená: `#86D295`
- Žlutá: `#FFCD00`
- Oranžová: `#F68D2E`
- Světle modrá: `#4EC3E0`
- Fialová: `#A05EB5`
- Červená: `#E2000F`
- Povrch aplikace: `#F5F6F4`
- Neutrální okraj: `#D8DDD9`

Font Koop nebyl v podkladech přiložen jako licencovaný soubor, proto je v souladu s manuálem použit systémový Arial. Logo bylo převzato přímo z dodaného vizuálního manuálu a není upravováno.

## Komponenty
Aplikační shell, levá navigace, horní lišta, globální hledání, metriky, grafy, tabulka případů, stavové štítky, produktové štítky, SLA časovač, průvodce nové škody, detail případu, časová osa, checklist, dokumentová tabulka, finanční karty, fraud indikátory, auditní log a responzivní modální okna.

## Role a oprávnění
Prototyp obsahuje výběr rolí: operátor, likvidátor, seniorní likvidátor, vedoucí, technik, lékařský posuzovatel, fraud specialista, právník, finanční pracovník, schvalovatel, klientské centrum, externí partner, auditor a administrátor. Produkční verze musí vynucovat RBAC/ABAC na backendu, nikoli pouze v UI.

## Workflow
Koncept → Přijato → Registrováno → Počáteční kontrola → Čeká na dokumenty → Posouzení krytí → Prohlídka / odborné posouzení / vyšetřování → Připraveno k rozhodnutí → Čeká na schválení → Připraveno k platbě → Zaplaceno → Uzavřeno. Dále jsou podporovány větve zamítnutí, regrese, soudního řízení, pozastavení a znovuotevření.

## Databázové schéma pro produkci
Doporučené tabulky: users, roles, permissions, teams, claims, claim_status_history, policies, policy_snapshots, coverages, participants, organisations, insured_objects, damage_items, injuries, documents, document_versions, tasks, workflows, communications, inspections, expert_reports, reserves, reserve_history, payments, payment_approvals, recoveries, fraud_indicators, investigations, legal_cases, complaints, partners, partner_assignments, catastrophe_events, audit_events, ai_recommendations, file_access_requests, quality_reviews.

Každá doménová tabulka má UUID, timestamps, autora změny a verzování. Finanční a auditní záznamy jsou append-only. Dokumenty jsou ukládány do objektového úložiště, metadata do PostgreSQL.

## Integrační adaptéry
Core systém (neutrální placeholder), systém smluv, klientský systém, dokumentové úložiště, účetnictví, platby, SSO, e-mail/SMS, datová schránka, mapy, databáze vozidel, partnerské sítě, kontrola pravosti dokumentů a analytický sklad.

## Bezpečnostní checklist
SSO + MFA, RBAC/ABAC, čtyři oči, šifrování, maskování dat, session timeout, antivirová kontrola, audit exportů a tisků, break-glass s důvodem, oddělení prostředí, tajemství mimo zdrojový kód, zálohy, DR testy, monitoring, rate limiting, CSP, ochrana API a pravidelná kontrola podle OWASP ASVS. Produkční návrh má respektovat GDPR a DORA.

## UX rozhodnutí
Rozhraní upřednostňuje rychlost práce a čitelnost před efekty. Klíčové informace jsou dostupné v souhrnném panelu. Další krok je vždy jasně viditelný. Barva není jediným nositelem stavu. AI návrhy mají zdroj, jistotu a vždy vyžadují lidské potvrzení.

## Mockované části
SSO, databáze, API, OCR, antivirová kontrola, elektronický podpis, datová schránka, platby, mapa, offline synchronizace, automatické workflow, externí integrace a AI model.

## Další vývoj
Rozdělit frontend na doménové moduly, doplnit backend (např. NestJS/Java/.NET), PostgreSQL, objektové úložiště, event store, background jobs, vyhledávání, e2e testy, autentizaci, auditní append-only úložiště, konfigurovatelný workflow engine a bezpečnostní penetrační testy.
