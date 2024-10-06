# iLLM-backend: Speak to LLM with Context

Entire iLLM backend

## Local Development Support

Run ollama as a single container

```shell
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
```

Install air using golang

```shell
go install github.com/air-verse/air@latest
```

Run server on local

```shell
air
```

If you don't want to use air

```shell
go mod download # Download the dependences
go run main.go #  run the server
```

## Supported

Context Extraction from **docx**. **pdf**, **text** file & **plain text** is supported

## Demo

Sample context:

```txt
"The short story is a crafted form in its own right. Short stories make use of plot, resonance and other dynamic components as in a novel, but typically to a lesser degree. While the short story is largely distinct from the novel or novella/short novel, authors generally draw from a common pool of literary techniques.[citation needed] The short story is sometimes referred to as a genre.[2]Determining what exactly defines a short story remains problematic.[3] A classic definition of a short story is that one should be able to read it in one sitting, a point most notably made in Edgar Allan Poe's essayThe Philosophy of Composition (1846).[4] H. G. Wells described the purpose of the short story as The jolly art, of making something very bright and moving; it may be horrible or pathetic or funny or profoundly illuminating, having only this essential, that it should take from fifteen to fifty minutes to read aloud.[5] According to William Faulkner, a short story is character-driven and a writer's job is to ...trot along behind him with a paper and pencil trying to keep up long enough to put down what he says and does.[6] Some authors have argued that a short story must have a strict form. Somerset Maugham thought that the short story must have a definite design, which includes a point of departure, a climax and a point of test; in other words, it must have a plot.[5] Hugh Walpole had a similar view: A story should be a story; a record of things happening full of incidents, swift movements, unexpected development, leading through suspense to a climax and a satisfying denouement.[5] This view of the short story as a finished product of art is however opposed by Anton Chekhov, who thought that a story should have neither a beginning nor an end. It should just be a slice of life, presented suggestively. In his stories, Chekhov does not round off the end but leaves it to the readers to draw their own conclusions.[5]Sukumar Azhikode defined a short story as a brief prose narrative with an intense episodic or anecdotal effect.[3] Flannery O'Connor emphasized the need to consider what is exactly meant by the descriptor short.[7] Short story writers may define their works as part of the artistic and personal expression of the form. They may also attempt to resist categorization by genre and fixed formation.[5] William Boyd, a British author and short story writer, has said: [a short story] seem[s] to answer something very deep in our nature as if, for the duration of its telling, something special has been created, some essence of our experience extrapolated, some temporary sense has been made of our common, turbulent journey towards the grave and oblivion.[8] In the 1880s, the term short story acquired its modern meaning – having initially referred to children's tales.[9] During the early to mid-20th century, the short story underwent expansive experimentation which further hindered attempts to comprehensively provide a definition.[3] Longer stories that cannot be called novels are sometimes considered novellas or novelettes and, like short stories, may be collected into the more marketable form of collections, of stories previously unpublished or published, but elsewhere.[citation needed] Sometimes, authors who do not have the time or money to write a novella or novel decide to write short stories instead, working out a deal with a popular website or magazine to publish them for profit.[citation needed] Around the world, the modern short story is comparable to lyrics, dramas, novels and essays – although examination of it as a major literary form remains diminished.[3][10]Length In terms of length, word count is typically anywhere from 1,000 to 4,000 for short stories; however, some works classified as short stories have up to 15,000 words. Stories of fewer than 1,000 words are sometimes referred to as short short stories, or flash fiction.[11] Short stories have no set length. In terms of word count, there is no official demarcation between an anecdote, a short story, and a novel. Rather, the form's parameters are given by the rhetorical and practical context in which a given story is produced and considered so that what constitutes a short story may differ between genres, countries, eras, and commentators.[12] Like the novel, the short story's predominant shape reflects the demands of the available markets for publication, and the evolution of the form seems closely tied to the evolution of the publishing industry and the submission guidelines of its constituent houses.[13] As a point of reference for the genre writer, the Science Fiction and Fantasy Writers of America define short story length in the Nebula Awards for science fiction submission guidelines as having fewer than 7,500 words"
```

Ask API: speak to LLM

![1725367238364](image/README/1725367238364.png)

Context API: provide context via pdf, docx, txt & plain text

![1725367275037](image/README/1725367275037.png)

## API Docs

```md
http:localhost:8090/swagger/index.html
```
