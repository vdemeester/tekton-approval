import { HttpClientModule } from '@angular/common/http';
import { MatButtonModule, MatExpansionModule, MatIconModule, MatTableModule } from '@angular/material';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ApprovalsService } from 'src/app/approvals.service';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { TektonApprovalsComponent } from './tekton-approvals/tekton-approvals.component';

@NgModule({
  declarations: [
    AppComponent,
    TektonApprovalsComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatTableModule,
    MatButtonModule,
    MatIconModule,
    MatExpansionModule,
    HttpClientModule
  ],
  providers: [ApprovalsService],
  bootstrap: [AppComponent]
})
export class AppModule { }
