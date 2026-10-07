# Counter

Small HTTP fixture for sandbox snapshot tests.

- `GET /readyz`: readiness; does not increment counters.
- `POST /`: increments memory and file counters; returns `{"memory":1,"file":1}`.
- Mount a writable directory at `/data` for the file counter.
- Listens on port 80.

Run **Actions → Publish counter → Run workflow** to publish AMD64/ARM64 images
to `ghcr.io/renato0307/counter`. The workflow uses GitHub's built-in token—no PAT.
Use the digest from the workflow summary. Make the GHCR package **public** for
anonymous image pulls.
