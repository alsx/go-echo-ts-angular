import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FbSigninComponent } from './fb-signin.component';

describe('FbSigninComponent', () => {
  let component: FbSigninComponent;
  let fixture: ComponentFixture<FbSigninComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FbSigninComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FbSigninComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
