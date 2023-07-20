# OCR API Service

This module provides an OCR (Optical Character Recognition) API service that detects and recognizes text in a PDF document. It utilizes the Google Cloud Vision API for OCR processing and provides a RESTful API for clients to interact with.

## API Endpoints
The OCR API service provides the following API endpoint:

POST `/ocr`: Accepts a PDF document as the request body and performs OCR to extract the text. It returns the recognized text along with the extracted paragraphs, headers, and the text itself.

## Examples
### Perform OCR on a PDF Document
#### Request:

```bash
curl -X POST -H "Content-Type: application/pdf" --data-binary @document.pdf http://localhost:8080/ocr

```


#### Response

```bash
{
  "title": "Title of the document",
  "abstract": "Abstract of the document",
  "main": {
    "paragraph1": "First paragraph...",
    "paragraph2": "Second paragraph...",
    ...
  },
  "references": "References of the document"
}

```