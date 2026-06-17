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

Repozitář je připravený pro automatické nasazení přes GitHub Actions. Po nahrání všech zdrojových souborů otevřete **Settings → Pages** a jako zdroj vyberte **GitHub Actions**.

## Technologie

React, TypeScript, Vite, Recharts a lokální perzistence přes `localStorage`.
