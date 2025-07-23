# 🎸 Groupie Tracker

**Groupie Tracker** is a responsive and interactive web application built with Go. It allows users to discover music bands and artists, view their members, formation dates, first albums, and explore upcoming concert locations and dates. All data is fetched live from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/).

---

## 🌟 Key Features

- 🔍 Search artists by name, creation year, or first album
- 🧑‍🎤 View artist profiles including members and images
- 📅 Display concert dates and locations
- 📱 Fully responsive and animated user interface
- ⚡ Real-time data fetched from the public API

---

## 🧱 Project Structure

```
.
├── main.go              # Main server setup and route handling
├── go.mod               # Go module configuration
├── /internal
│   ├── fetch.go         # API data fetching functions
│   ├── struct.go        # Data models (Artist, Relation, etc.)
│   └── search.go        # Optional: logic for filtering results
├── /templates
│   ├── index.html       # Homepage with artist cards
│   └── details.html     # Detailed artist profile page
└── /static
    ├── style1.css       # Homepage styling
    └── details.css      # Artist details page styling
```

---

## 🧠 How It Works

1. The homepage fetches and displays all artists as styled cards.
2. Clicking on a card redirects to the artist’s detail page.
3. Each detail page shows:
   - Band members  
   - First album date  
   - Tour locations and dates  
4. Pages are rendered using Go’s built-in `html/template` engine.

---

## 📡 API Integration

- `GET /api/artists` → retrieves a list of all artists
- `GET /api/relation/{id}` → retrieves tour locations and dates for the selected artist
- Data is fetched using Go’s `net/http` and parsed via `encoding/json`

---

## 🚀 How to Run

### Prerequisites

- Go 1.18 or newer must be installed
- Internet connection (required to fetch API data)

### Instructions

```bash
git clone https://github.com/yourusername/groupie-tracker.git
cd groupie-tracker
go run main.go
```

Open your browser and navigate to:  
👉 `http://localhost:8080`

---

## 📌 Notes

- The search engine (`search.go`) is simple and can be expanded (e.g., search by members or album)
- JSON links or raw API values are displayed in some cases for demonstration purposes
- Artist ID 21 is excluded manually for testing purposes

---

## 🧑‍💻 Developer

This project was built by **Adel Kanaan** as part of a learning journey in:

- Backend development with Go  
- RESTful API consumption  
- JSON parsing and data binding  
- Server-side rendering with HTML templates  
- Responsive web design with modern CSS

> 🌐 Portfolio: [https://noorkanaan.github.io/NoorKanaan.com/](https://noorkanaan.github.io/NoorKanaan.com/)

---

## 📄 License

This project is open-source and intended for educational and demonstrational purposes only.
