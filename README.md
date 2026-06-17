# KLIK – Kooperativní likvidační informační komplex

Demonstrační projekt vytvořený **Danielem Třetinou** pro účely výběrového řízení.

> Toto není oficiální webová stránka ani interní systém společnosti Kooperativa. Aplikace vznikla pouze jako autorský demonstrační projekt na základě veřejně dostupných informací a obecných principů likvidace pojistných událostí. Všechna data jsou syntetická.

## Lokální spuštění

```bash
npm ci
npm run dev
```

## Produkční build

```bash
npm ci
npm run build
```

## GitHub Pages

Repozitář obsahuje workflow `.github/workflows/deploy-pages.yml`. V nastavení repozitáře otevřete **Settings → Pages** a jako zdroj vyberte **GitHub Actions**. Po pushi do větve `main` se aplikace automaticky sestaví a nasadí.

## Technologie

React, TypeScript, Vite, Recharts a lokální perzistence přes `localStorage`.
