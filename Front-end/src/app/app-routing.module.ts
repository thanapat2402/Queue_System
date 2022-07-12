import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { WebAdminComponent } from './components/web-admin/web-admin.component';

const routes: Routes = [
  { path: '', redirectTo: 'admin', pathMatch: 'full' },
  { path: 'admin', component: WebAdminComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
