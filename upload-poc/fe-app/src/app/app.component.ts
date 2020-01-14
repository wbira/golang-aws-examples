import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { pluck, mergeMap, take } from 'rxjs/operators';

interface PresignedUrl {
  url: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private http: HttpClient) {
  }

  // TODO CHANGE API URL!!!
  private apiUrl = "https://3gv0w614zk.execute-api.eu-west-1.amazonaws.com/Stage/hello"

  public onFileUpload(event: Event): void {
    const file = event.target['files'][0]
    this.uploadFile(file)
  }

  private uploadFile(file: File) {
    this.getPresignedUrls(file)
      .pipe(
        mergeMap((url: string) =>
          this.postFileData(url, file)
        ), take(1)).subscribe()
  }

  private getPresignedUrls({ type, name }: File): Observable<string> {
    return this.http.get<PresignedUrl>(`${this.apiUrl}?contentType=${type}&name=${name}`).pipe(pluck('url'))
  }

  private postFileData(url: string, file: File): Observable<any> {
    let headers = new HttpHeaders();
    headers = headers.append('Content-Type', file.type);
    return this.http.put(url, file, { headers })
  }
}
