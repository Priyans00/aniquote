

```markdown
# AniQuote - Anime Quotes API

A simple REST API built in Go that serves anime-related quotes. This API provides random anime quotes, quotes by ID, character, or anime name.

```

The API will be available at `https://aniquote-production.up.railway.app`


## 🌐 API Endpoints

### Get Quotes
- **GET** `/api/quote`
  - Returns a random quote 
  - Example: `/api/quote`

**Response:**
```json
[
    {
        "no": 1138,
        "anime": "(One Piece)",
        "character": "Brook",
        "quote": "What do you know of death?!! Have you ever died?!! You think death will preserve your cause forever?!! Ridiculous!! Death leaves nothing behind!! Once a person passes on, nothing remains but dead bones!! If there is one thing I can’t stand, it is a person with no respect for life!! \r\n"
    }
]
```

### Get Quote by ID
- **GET** `/api/quotes/:id`
  - Returns a specific quote by its ID
  - Example: `/api/quotes/15`

**Response:**
```json
[
    {
        "no": 15,
        "anime": "(Death Parade)",
        "character": "Nona",
        "quote": "Everyone Makes Mistakes. But people's feelings can manifest through the most subtle of expressions."
    }
]
```

### Get Quotes by Character Name
- **GET** `/api/quotes/character/:name`
  - Returns all quotes by a specific character
  - Example: `/api/quotes/character/Kamina`

**Response:**
```json
[
    {
        "no": 75,
        "anime": "(Gurren Lagann)",
        "character": "Kamina",
        "quote": "Don't believe in the me that believe in you, don't believe in the you that believes in me, believe in yourself who believes in you! "
    },
    {
        "no": 2106,
        "anime": "(Tengen Toppa Gurren Laggan)",
        "character": "Kamina",
        "quote": "Don't be distracted by the what-if's, should-have's, and if-only's. The one thing you choose for yourself - that is the truth of your universe."
    },.......

]
```

### Get Quotes by Anime Name
- **GET** `/api/quotes/anime/:name`
  - Returns all quotes from a specific anime
  - Example: `/api/quotes/anime/Berserk`

**Response:**
```json
[
    {
        "no": 2661,
        "anime": "(Berserk)",
        "character": "Berserk",
        "quote": "Providence may guide a man to meet one specific person, even if such guidance eventually leads him to darkness. Man simply cannot forsake the beauty of his own chosen path. When will man learn a way to control his soul?"
    },
    {
        "no": 2662,
        "anime": "(Berserk)",
        "character": "Berserk",
        "quote": "Dreams, ambition, love, hope; in this world, could the glories of a youthful heart be.. forbidden?"
    },.......
]
```

## 🚧 Error Responses
```json
{
  "error": "Quote not found"
}
```

## NOTE:
Api rate limit is 10 requests per 10 sec.

Quote fetching using Character and Anime name is case sensitive and contains the entire name of the season . Its still in development.

## 🤝 Contributing
Contributions are welcome! Please open an issue first to discuss what you'd like to change, or submit a pull request directly.

## 📄 License
[MIT](https://choosealicense.com/licenses/mit/) 

---

⭐ Star this repository on [GitHub](https://github.com/Priyans00/aniquote)
```

